#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# Following script is adapted from https://github.com/ava-labs/subnet-evm/blob/master/scripts/install_avalanchego_release.sh
set -e

# Load the versions
TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)
source "$TELEPORTER_PATH"/scripts/versions.sh
source "$TELEPORTER_PATH"/scripts/constants.sh

############################
# download subnet-evm
# https://github.com/ava-labs/subnet-evm/releases
GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)
BASEDIR=${BASEDIR:-"/tmp/subnet-evm-release"}
SUBNET_EVM_BUILD_PATH=${SUBNET_EVM_BUILD_PATH:-${BASEDIR}/subnet-evm}

mkdir -p ${BASEDIR}

SUBNET_EVM_DOWNLOAD_URL=https://github.com/ava-labs/subnet-evm/releases/download/${SUBNET_EVM_VERSION}/subnet-evm_${SUBNET_EVM_VERSION#v}_linux_${GOARCH}.tar.gz
SUBNET_EVM_DOWNLOAD_PATH=${BASEDIR}/subnet-evm-linux-${GOARCH}-${SUBNET_EVM_VERSION}.tar.gz

if [[ ${GOOS} == "darwin" ]]; then
  SUBNET_EVM_DOWNLOAD_URL=https://github.com/ava-labs/subnet-evm/releases/download/${SUBNET_EVM_VERSION}/subnet-evm_${SUBNET_EVM_VERSION#v}_darwin_${GOARCH}.tar.gz
  SUBNET_EVM_DOWNLOAD_PATH=${BASEDIR}/subnet-evm-darwin-${GOARCH}-${SUBNET_EVM_VERSION}.tar.gz
fi

BUILD_DIR=${SUBNET_EVM_BUILD_PATH}-${SUBNET_EVM_VERSION}

extract_archive() {
  mkdir -p ${BUILD_DIR}

  if [[ ${SUBNET_EVM_DOWNLOAD_PATH} == *.tar.gz ]]; then
    tar xzvf ${SUBNET_EVM_DOWNLOAD_PATH} --directory ${BUILD_DIR}
  elif [[ ${SUBNET_EVM_DOWNLOAD_PATH} == *.zip ]]; then
    unzip ${SUBNET_EVM_DOWNLOAD_PATH} -d ${BUILD_DIR}
    mv ${BUILD_DIR}/build/* ${BUILD_DIR}
    rm -rf ${BUILD_DIR}/build/
  fi
}

# first check if we already have the archive
if [[ -f ${SUBNET_EVM_DOWNLOAD_PATH} ]]; then
  # if the download path already exists, extract and exit
  echo "found subnet-evm ${SUBNET_EVM_VERSION} at ${SUBNET_EVM_DOWNLOAD_PATH}"

  extract_archive
else
  # try to download the archive if it exists
  if curl -s --head --request GET ${SUBNET_EVM_DOWNLOAD_URL} | grep "302" > /dev/null; then
    echo "${SUBNET_EVM_DOWNLOAD_URL} found"
    echo "downloading to ${SUBNET_EVM_DOWNLOAD_PATH}"
    curl -L ${SUBNET_EVM_DOWNLOAD_URL} -o ${SUBNET_EVM_DOWNLOAD_PATH}

    extract_archive
  else
    # else the version is a git commit (or it's invalid)
    GIT_CLONE_URL=https://github.com/ava-labs/subnet-evm.git
    GIT_CLONE_PATH=${BASEDIR}/subnet-evm-repo/
    
    # check to see if the repo already exists, if not clone it 
    if [[ ! -d ${GIT_CLONE_PATH} ]]; then
      echo "cloning ${GIT_CLONE_URL} to ${GIT_CLONE_PATH}"
      git clone --no-checkout ${GIT_CLONE_URL} ${GIT_CLONE_PATH}
    fi

    # check to see if the commitish exists in the repo
    WORKDIR=$(pwd)

    cd ${GIT_CLONE_PATH}

    git fetch

    echo "checking out ${SUBNET_EVM_VERSION}"

    set +e
    # try to checkout the branch
    git checkout origin/${SUBNET_EVM_VERSION} > /dev/null 2>&1
    CHECKOUT_STATUS=$?
    set -e

    # if it's not a branch, try to checkout the commit 
    if [[ $CHECKOUT_STATUS -ne 0 ]]; then
      set +e
      git checkout ${SUBNET_EVM_VERSION} > /dev/null 2>&1
      CHECKOUT_STATUS=$?
      set -e

      if [[ $CHECKOUT_STATUS -ne 0 ]]; then
        echo
        echo "'${VERSION}' is not a valid release tag, commit hash, or branch name"
        exit 1
      fi
    fi

    COMMIT=$(git rev-parse HEAD)

    # use the commit hash instead of the branch name or tag
    BUILD_DIR=${SUBNET_EVM_BUILD_PATH}-${COMMIT}

    # if the build-directory doesn't exist, build subnet-evm
    if [[ ! -d ${BUILD_DIR} ]]; then    
      echo "building subnet-evm ${COMMIT} to ${BUILD_DIR}"
      ./scripts/build.sh
      mkdir -p ${BUILD_DIR}

      mv ${GIT_CLONE_PATH}/build/* ${BUILD_DIR}/
    fi

    cd $WORKDIR
  fi
fi

SUBNET_EVM_PATH=${SUBNET_EVM_BUILD_PATH}/subnet-evm
mkdir -p ${SUBNET_EVM_BUILD_PATH}
    
cp ${BUILD_DIR}/subnet-evm ${SUBNET_EVM_PATH}


echo "Installed Subnet-EVM release ${SUBNET_EVM_VERSION}"
echo "Subnet-EVM Path: ${SUBNET_EVM_PATH}"