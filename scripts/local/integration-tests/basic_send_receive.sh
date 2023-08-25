#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e # Stop on first error

# Variables provided by run_setup.sh:
#   c_chain_url
#   user_private_key
#   user_address_bytes
#   user_address
#   relayer_address
#   subnet_a_chain_id
#   subnet_b_chain_id
#   subnet_a_subnet_id
#   subnet_b_subnet_id
#   subnet_a_url
#   subnet_b_url
#   subnet_a_chain_id_hex
#   subnet_b_chain_id_hex
#   subnet_a_subnet_id_hex
#   subnet_b_subnet_id_hex
#   teleporter_contract_address
#   warp_messenger_precompile_addr

# Test covers:
# - Sending bidirectional cross chain messages between two chains, by calling Teleporter contract sendCrossChainMessage function directly.
# - Checking message delivery for message that was sent on destination chain.
# - Calling to the warp precompile address directly for blockchain ID.

# Deploy a test ERC20 to be used in the E2E test.
cd contracts
erc20_deploy_result=$(forge create --private-key $user_private_key src/Mocks/ExampleERC20.sol:ExampleERC20 --rpc-url $subnet_a_url)
erc20_contract_address_a=$(parseContractAddress "$erc20_deploy_result")
echo "Test ERC20 contract deployed to $erc20_contract_address_a on subnet A"
erc20_deploy_result=$(forge create --private-key $user_private_key src/Mocks/ExampleERC20.sol:ExampleERC20 --rpc-url $subnet_b_url)
erc20_contract_address_b=$(parseContractAddress "$erc20_deploy_result")
echo "Test ERC20 contract deployed to $erc20_contract_address_b on subnet B"

###
# Send from subnet A -> subnet B
###
echo "Sending from subnet A to subnet B"
blockchainID=$(cast call $warp_messenger_precompile_addr "getBlockchainID()(bytes32)" --rpc-url $subnet_a_url)
echo "Got blockchain ID $blockchainID"

echo "Sending call to teleporter contract address $teleporter_contract_address $subnet_b_chain_id_hex $subnet_a_url"
result=$(cast call $teleporter_contract_address "getNextMessageID(bytes32)(uint256)" $subnet_b_chain_id_hex --rpc-url $subnet_a_url)
echo "Next message ID for subnet $subnet_b_chain_id_hex is $result"

# Directly send a transaction to the teleporter contract sendCrossChainMessage function.
send_cross_subnet_message_destination_chain_id=$subnet_b_chain_id_hex
send_cross_subnet_message_destination_address=abcedf1234abcedf1234abcedf1234abcedf1234
send_bytes32=000000000000000000000000$send_cross_subnet_message_destination_address
send_cross_subnet_message_fee_amount=00000000000000000000000000000000000000000000000000000000000000FF
send_cross_subnet_message_required_gas_limit=0000000000000000000000000000000000000000000000000000000000001000
send_cross_subnet_message_message_data=cafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabecafe

# Approve the Teleporter contract to some ERC20 tokens from the user account we're using to send transactions
cast send $erc20_contract_address_a "approve(address,uint256)(bool)" $teleporter_contract_address \
    000000000000000000000000000000000000000000FFFFFFFFFFFFFFFFFFFFFF \
    --private-key $user_private_key --rpc-url $subnet_a_url
result=$(cast call $erc20_contract_address_a "allowance(address,address)(uint256)" $user_address $teleporter_contract_address --rpc-url $subnet_a_url)
if [[ $result != 309485009821345068724781055 ]]; then # FFFFFFFFFFFFFFFFFFFFFF in decimal form is 309485009821345068724781055
    echo $result
    echo "Error approving Teleporter contract to spend ERC20 from user account."
    exit 1
fi

echo "Approved the Teleporter contract to spend the test ERC20 token from the user account."

startID=$(cast call $teleporter_contract_address "sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes))(uint256)" "($send_cross_subnet_message_destination_chain_id,$send_cross_subnet_message_destination_address,($erc20_contract_address_a,$send_cross_subnet_message_fee_amount),$send_cross_subnet_message_required_gas_limit,[],$send_cross_subnet_message_message_data)" --from $user_address --rpc-url $subnet_a_url)
echo "Got starting ID $startID to teleport address $teleporter_contract_address"
echo "Got Ids $subnet_a_chain_id_hex $subnet_b_chain_id_hex $subnet_a_subnet_id $subnet_b_subnet_id"
cast send $teleporter_contract_address "sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes))(uint256)" "($send_cross_subnet_message_destination_chain_id,$send_cross_subnet_message_destination_address,($erc20_contract_address_a,$send_cross_subnet_message_fee_amount),$send_cross_subnet_message_required_gas_limit,[],$send_cross_subnet_message_message_data)" --private-key $user_private_key --rpc-url $subnet_a_url

retry_count=0
received=$(cast call $teleporter_contract_address "messageReceived(bytes32,uint256)(bool)" $subnet_a_chain_id_hex $startID --rpc-url $subnet_b_url)
until [[ $received == "true" ]]
do
    if [[ retry_count -ge 10 ]]; then
        echo "Destination chain on subnet B did not receive message before timeout."
        exit 1
    fi
    echo "Waiting for destination chain on subnet B to receive message ID $startID. Retry count: $retry_count"
    sleep 3

    received=$(cast call $teleporter_contract_address "messageReceived(bytes32,uint256)(bool)" $subnet_a_chain_id_hex $startID --rpc-url $subnet_b_url)
    retry_count=$((retry_count+1))
done

echo "Received on subnet B is $received"


###
# Send from subnet B -> subnet A
###
echo "Sending from subnet B to subnet A"
blockchainID=$(cast call $warp_messenger_precompile_addr "getBlockchainID()(bytes32)" --rpc-url $subnet_b_url)
echo "Got blockchain ID $blockchainID"

echo "Sending call to teleporter contract address $teleporter_contract_address $subnet_a_chain_id_hex $subnet_b_url"
result=$(cast call $teleporter_contract_address "getNextMessageID(bytes32)(uint256)" $subnet_a_chain_id_hex --rpc-url $subnet_b_url)
echo "Next message ID for subnet $subnet_a_chain_id_hex is $result"

# Directly send a few transaction to the teleporter contract sendCrossChainMessage function.
send_cross_subnet_message_destination_chain_id=$subnet_a_chain_id_hex
send_cross_subnet_message_destination_address=abcedf1234abcedf1234abcedf1234abcedf1234
send_bytes32=000000000000000000000000$send_cross_subnet_message_destination_address
send_cross_subnet_message_fee_amount=00000000000000000000000000000000000000000000000000000000000000FF
send_cross_subnet_message_required_gas_limit=0000000000000000000000000000000000000000000000000000000000001000
send_cross_subnet_message_message_data=cafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabecafe

# Approve the teleporter contract to some ERC20 tokens from the user account we're using to send transactions
cast send $erc20_contract_address_b "approve(address,uint256)(bool)" $teleporter_contract_address 000000000000000000000000000000000000000000FFFFFFFFFFFFFFFFFFFFFF --private-key $user_private_key --rpc-url $subnet_b_url
result=$(cast call $erc20_contract_address_b "allowance(address,address)(uint256)" $user_address $teleporter_contract_address --rpc-url $subnet_b_url)
if [[ $result != 309485009821345068724781055 ]]; then # FFFFFFFFFFFFFFFFFFFFFF in decimal form is 309485009821345068724781055
    echo $result
    echo "Error approving Teleporter contract to spend ERC20 from user account."
    exit 1
fi

echo "Approved the Teleporter contract to spend the test ERC20 token from the user account."

startID=$(cast call $teleporter_contract_address "sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes))(uint256)" "($send_cross_subnet_message_destination_chain_id,$send_cross_subnet_message_destination_address,($erc20_contract_address_b,$send_cross_subnet_message_fee_amount),$send_cross_subnet_message_required_gas_limit,[],$send_cross_subnet_message_message_data)" --from $user_address --rpc-url $subnet_b_url)
echo "Got starting ID $startID to teleport address $teleporter_contract_address"
echo "Got Ids $subnet_b_chain_id_hex $subnet_a_chain_id_hex $subnet_b_subnet_id $subnet_a_subnet_id"
cast send $teleporter_contract_address "sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes))(uint256)" "($send_cross_subnet_message_destination_chain_id,$send_cross_subnet_message_destination_address,($erc20_contract_address_b,$send_cross_subnet_message_fee_amount),$send_cross_subnet_message_required_gas_limit,[],$send_cross_subnet_message_message_data)" --private-key $user_private_key --rpc-url $subnet_b_url

retry_count=0
received=$(cast call $teleporter_contract_address "messageReceived(bytes32,uint256)(bool)" $subnet_b_chain_id_hex $startID --rpc-url $subnet_a_url)
until [[ $received == "true" ]]
do
    if [[ retry_count -ge 10 ]]; then
        echo "Destination chain on subnet A did not receive message before timeout."
        exit 1
    fi
    echo "Waiting for destination on subnet A chain to receive message ID $startID. Retry count: $retry_count"
    sleep 3

    received=$(cast call $teleporter_contract_address "messageReceived(bytes32,uint256)(bool)" $subnet_b_chain_id_hex $startID  --rpc-url $subnet_a_url)
    retry_count=$((retry_count+1))
done

echo "Received on subnet A is $received"
