function createLibrary {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/library.libBadge.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/library.libBadge.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-library --tls.certfiles ${PWD}/organizations/fabric-ca/library/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-library.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-library.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-library.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-library.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/library.libBadge.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	fabric-ca-client register --caname ca-library --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/library/tls-cert.pem
  set +x

  echo
	echo "Register peer1"
  echo
  set -x
	fabric-ca-client register --caname ca-library --id.name peer1 --id.secret peer1pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/library/tls-cert.pem
  set +x

  echo
  echo "Register user"
  echo
  set -x
  fabric-ca-client register --caname ca-library --id.name user1 --id.secret user1pw --id.type client --tls.certfiles ${PWD}/organizations/fabric-ca/library/tls-cert.pem
  set +x

  echo
  echo "Register the org admin"
  echo
  set -x
  fabric-ca-client register --caname ca-library --id.name libraryadmin --id.secret libraryadminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/library/tls-cert.pem
  set +x

	mkdir -p organizations/peerOrganizations/library.libBadge.com/peers
  mkdir -p organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com
  mkdir -p organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-library -M ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/msp --csr.hosts peer0.library.libBadge.com --tls.certfiles ${PWD}/organizations/fabric-ca/library/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-library -M ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/tls --enrollment.profile tls --csr.hosts peer0.library.libBadge.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/library/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/tls/server.key

  mkdir -p ${PWD}/organizations/peerOrganizations/library.libBadge.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/library.libBadge.com/msp/tlscacerts/ca.crt

  mkdir -p ${PWD}/organizations/peerOrganizations/library.libBadge.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/library.libBadge.com/tlsca/tlsca.library.libBadge.com-cert.pem

  mkdir -p ${PWD}/organizations/peerOrganizations/library.libBadge.com/ca
  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer0.library.libBadge.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/library.libBadge.com/ca/ca.library.libBadge.com-cert.pem


  echo
  echo "## Generate the peer1 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer1:peer1pw@localhost:7054 --caname ca-library -M ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/msp --csr.hosts peer1.library.libBadge.com --tls.certfiles ${PWD}/organizations/fabric-ca/library/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/msp/config.yaml

  echo
  echo "## Generate the peer1-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:7054 --caname ca-library -M ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/tls --enrollment.profile tls --csr.hosts peer1.library.libBadge.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/library/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/tls/server.key

  mkdir -p ${PWD}/organizations/peerOrganizations/library.libBadge.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/library.libBadge.com/msp/tlscacerts/ca.crt

  mkdir -p ${PWD}/organizations/peerOrganizations/library.libBadge.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/library.libBadge.com/tlsca/tlsca.library.libBadge.com-cert.pem

  mkdir -p ${PWD}/organizations/peerOrganizations/library.libBadge.com/ca
  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/peers/peer1.library.libBadge.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/library.libBadge.com/ca/ca.library.libBadge.com-cert.pem



  mkdir -p organizations/peerOrganizations/library.libBadge.com/users
  mkdir -p organizations/peerOrganizations/library.libBadge.com/users/User1@library.libBadge.com

  echo
  echo "## Generate the user msp"
  echo
  set -x
	fabric-ca-client enroll -u https://user1:user1pw@localhost:7054 --caname ca-library -M ${PWD}/organizations/peerOrganizations/library.libBadge.com/users/User1@library.libBadge.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/library/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/library.libBadge.com/users/User1@library.libBadge.com/msp/config.yaml

  mkdir -p organizations/peerOrganizations/library.libBadge.com/users/Admin@library.libBadge.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://libraryadmin:libraryadminpw@localhost:7054 --caname ca-library -M ${PWD}/organizations/peerOrganizations/library.libBadge.com/users/Admin@library.libBadge.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/library/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/library.libBadge.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/library.libBadge.com/users/Admin@library.libBadge.com/msp/config.yaml

}


function createStudent {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/student.libBadge.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/student.libBadge.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:8054 --caname ca-student --tls.certfiles ${PWD}/organizations/fabric-ca/student/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-student.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-student.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-student.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-student.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/student.libBadge.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	fabric-ca-client register --caname ca-student --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/student/tls-cert.pem
  set +x

  echo
	echo "Register peer1"
  echo
  set -x
	fabric-ca-client register --caname ca-student --id.name peer1 --id.secret peer1pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/student/tls-cert.pem
  set +x

  echo
  echo "Register user"
  echo
  set -x
  fabric-ca-client register --caname ca-student --id.name user1 --id.secret user1pw --id.type client --tls.certfiles ${PWD}/organizations/fabric-ca/student/tls-cert.pem
  set +x

  echo
  echo "Register the org admin"
  echo
  set -x
  fabric-ca-client register --caname ca-student --id.name studentadmin --id.secret studentadminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/student/tls-cert.pem
  set +x

	mkdir -p organizations/peerOrganizations/student.libBadge.com/peers
  mkdir -p organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com
  mkdir -p organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer0:peer0pw@localhost:8054 --caname ca-student -M ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/msp --csr.hosts peer0.student.libBadge.com --tls.certfiles ${PWD}/organizations/fabric-ca/student/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:8054 --caname ca-student -M ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/tls --enrollment.profile tls --csr.hosts peer0.student.libBadge.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/student/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/tls/server.key

  mkdir -p ${PWD}/organizations/peerOrganizations/student.libBadge.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/student.libBadge.com/msp/tlscacerts/ca.crt

  mkdir -p ${PWD}/organizations/peerOrganizations/student.libBadge.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/student.libBadge.com/tlsca/tlsca.student.libBadge.com-cert.pem

  mkdir -p ${PWD}/organizations/peerOrganizations/student.libBadge.com/ca
  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer0.student.libBadge.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/student.libBadge.com/ca/ca.student.libBadge.com-cert.pem

  echo
  echo "## Generate the peer1 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer1:peer1pw@localhost:8054 --caname ca-student -M ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/msp --csr.hosts peer1.student.libBadge.com --tls.certfiles ${PWD}/organizations/fabric-ca/student/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/msp/config.yaml

  echo
  echo "## Generate the peer1-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:8054 --caname ca-student -M ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/tls --enrollment.profile tls --csr.hosts peer1.student.libBadge.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/student/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/tls/server.key

  mkdir -p ${PWD}/organizations/peerOrganizations/student.libBadge.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/student.libBadge.com/msp/tlscacerts/ca.crt

  mkdir -p ${PWD}/organizations/peerOrganizations/student.libBadge.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/student.libBadge.com/tlsca/tlsca.student.libBadge.com-cert.pem

  mkdir -p ${PWD}/organizations/peerOrganizations/student.libBadge.com/ca
  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/peers/peer1.student.libBadge.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/student.libBadge.com/ca/ca.student.libBadge.com-cert.pem


  mkdir -p organizations/peerOrganizations/student.libBadge.com/users
  mkdir -p organizations/peerOrganizations/student.libBadge.com/users/User1@student.libBadge.com

  echo
  echo "## Generate the user msp"
  echo
  set -x
	fabric-ca-client enroll -u https://user1:user1pw@localhost:8054 --caname ca-student -M ${PWD}/organizations/peerOrganizations/student.libBadge.com/users/User1@student.libBadge.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/student/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/student.libBadge.com/users/User1@student.libBadge.com/msp/config.yaml

  mkdir -p organizations/peerOrganizations/student.libBadge.com/users/Admin@student.libBadge.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://studentadmin:studentadminpw@localhost:8054 --caname ca-student -M ${PWD}/organizations/peerOrganizations/student.libBadge.com/users/Admin@student.libBadge.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/student/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/student.libBadge.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/student.libBadge.com/users/Admin@student.libBadge.com/msp/config.yaml

}



function createOrderer {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/ordererOrganizations/libBadge.com

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/ordererOrganizations/libBadge.com
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:9054 --caname ca-orderer --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/ordererOrganizations/libBadge.com/msp/config.yaml


  echo
	echo "Register raft1"
  echo
  set -x
	fabric-ca-client register --caname ca-orderer --id.name raft1 --id.secret raft1pw --id.type orderer --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  echo
	echo "Register raft2"
  echo
  set -x
	fabric-ca-client register --caname ca-orderer --id.name raft2 --id.secret raft2pw --id.type orderer --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  echo
  echo "Register the orderer admin"
  echo
  set -x
  fabric-ca-client register --caname ca-orderer --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

	mkdir -p organizations/ordererOrganizations/libBadge.com/orderers
  # mkdir -p organizations/ordererOrganizations/libBadge.com/orderers/libBadge.com

  mkdir -p organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com
  mkdir -p organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com

  echo
  echo "## Generate the raft1 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://raft1:raft1pw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/msp --csr.hosts raft1.libBadge.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/msp/config.yaml

  echo
  echo "## Generate the raft1-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://raft1:raft1pw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/tls --enrollment.profile tls --csr.hosts raft1.libBadge.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/tls/ca.crt
  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/tls/signcerts/* ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/tls/server.crt
  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/tls/keystore/* ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/tls/server.key

  mkdir -p ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/msp/tlscacerts/tlsca.libBadge.com-cert.pem

  mkdir -p ${PWD}/organizations/ordererOrganizations/libBadge.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft1.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/libBadge.com/msp/tlscacerts/tlsca.libBadge.com-cert.pem

  echo
  echo "## Generate the raft2 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://raft2:raft2pw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/msp --csr.hosts raft2.libBadge.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/msp/config.yaml

  echo
  echo "## Generate the raft2-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://raft2:raft2pw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/tls --enrollment.profile tls --csr.hosts raft2.libBadge.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/tls/ca.crt
  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/tls/signcerts/* ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/tls/server.crt
  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/tls/keystore/* ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/tls/server.key

  mkdir -p ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/msp/tlscacerts/tlsca.libBadge.com-cert.pem

  mkdir -p ${PWD}/organizations/ordererOrganizations/libBadge.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/orderers/raft2.libBadge.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/libBadge.com/msp/tlscacerts/tlsca.libBadge.com-cert.pem



  mkdir -p organizations/ordererOrganizations/libBadge.com/users
  mkdir -p organizations/ordererOrganizations/libBadge.com/users/Admin@libBadge.com

  echo
  echo "## Generate the admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://ordererAdmin:ordererAdminpw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/libBadge.com/users/Admin@libBadge.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/libBadge.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/libBadge.com/users/Admin@libBadge.com/msp/config.yaml

}
