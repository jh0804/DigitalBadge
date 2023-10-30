#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a collection of bash functions used by different scripts

export CORE_PEER_TLS_ENABLED=true
# Orderer CA environment variables
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/msp/tlscacerts/tlsca.libBadge.com-cert.pem
export ORDERER2_CA=${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/orderer2.libBadge.com/msp/tlscacerts/tlsca.libBadge.com-cert.pem

# LibraryOrg environment variables
export PEER0_LIBRARY_CA=${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/tls/ca.crt
export PEER1_LIBRARY_CA=${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/tls/ca.crt

# StudentOrg environment variables
export PEER0_STUDENT_CA=${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/tls/ca.crt
export PEER1_STUDENT_CA=${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/tls/ca.crt


# Set OrdererOrg.Admin globals
setOrdererGlobals() {
  export CORE_PEER_LOCALMSPID="OrdererMSP"
  export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/msp/tlscacerts/tlsca.libBadge.com-cert.pem
  export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/ordererOrganizations/libBadge.com/users/Admin@libBadge.com/msp
}

# Set environment variables for the peer org
setGlobals() {
  local USING_ORG=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$1
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi
  
  echo "Using organization ${USING_ORG}"
  local MYPEER=$2
  
  if [ $USING_ORG == 'library' ] && [ $MYPEER -eq 0 ]; then
    export CORE_PEER_LOCALMSPID="LibraryMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_LIBRARY_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/library.libBadge.com/users/Admin@library.libBadge.com/msp
    export CORE_PEER_ADDRESS=localhost:7051
    export NUM="ch1"
  elif [ $USING_ORG == 'library' ] && [ $MYPEER -eq 1 ]; then
    export CORE_PEER_LOCALMSPID="LibraryMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER1_LIBRARY_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/library.libBadge.com/users/Admin@library.libBadge.com/msp
    export CORE_PEER_ADDRESS=localhost:8051
    export NUM="ch1"
  elif [ $USING_ORG == 'student' ] && [ $MYPEER -eq 0 ]; then
    export CORE_PEER_LOCALMSPID="StudentMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_STUDENT_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/student.libBadge.com/users/Admin@student.libBadge.com/msp
    export CORE_PEER_ADDRESS=localhost:9051 
    export NUM="ch1"
  elif [ $USING_ORG == 'student' ] && [ $MYPEER -eq 1 ]; then
    export CORE_PEER_LOCALMSPID="StudentMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER1_STUDENT_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/student.libBadge.com/users/Admin@student.libBadge.com/msp
    export CORE_PEER_ADDRESS=localhost:10051  
    export NUM="ch1"
  else
    echo "================== ERROR !!! ORG Unknown =================="
  fi

  if [ "$VERBOSE" == "true" ]; then
    env | grep CORE
  fi
}

# parsePeerConnectionParameters $@
# Helper function that sets the peer connection parameters for a chaincode
# operation
parsePeerConnectionParameters() {

  PEER_CONN_PARMS=""
  PEERS=""
  
  while [ "$#" -gt 0 ]; do
    setGlobals $1 $2
    PEER="peer$2.$1"
    ## Set peer adresses
    PEERS="$PEERS $PEER"
    PEER_CONN_PARMS="$PEER_CONN_PARMS --peerAddresses $CORE_PEER_ADDRESS"
    if [ $1 == 'library' ]; then
      ## Set path to TLS certificate
      TLSINFO=$(eval echo "--tlsRootCertFiles \$PEER$2_LIBRARY_CA")
    elif [ $1 == 'student' ]; then
      ## Set path to TLS certificate
      TLSINFO=$(eval echo "--tlsRootCertFiles \$PEER$2_STUDENT_CA")
    fi
    PEER_CONN_PARMS="$PEER_CONN_PARMS $TLSINFO"
    # shift by one to get to the next organization
    shift 2
  done
  # remove leading space for output
  PEERS="$(echo -e "$PEERS" | sed -e 's/^[[:space:]]*//')"
}

verifyResult() {
  if [ $1 -ne 0 ]; then
    echo $'\e[1;31m'!!!!!!!!!!!!!!! $2 !!!!!!!!!!!!!!!!$'\e[0m'
    echo
    exit 1
  fi
}
