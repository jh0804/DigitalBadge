# DigitalBadge

*네트워크 구축 순서*
export FABRIC_CFG_PATH=$PWD

1.crypto-config.yaml

./bin/cryptogen generate --config=./crypto-config.yaml

2.configtx.yaml

./bin/configtxgen -profile OrdererGenesis -outputBlock ./config/genesis.block -channelID channelbadge1
./bin/configtxgen -profile Channel1 -outputCreateChannelTx ./config/channel1.tx -channelID channelbadge1
./bin/configtxgen -profile Channel1 -outputAnchorPeersUpdate ./config/LibraryOrganchors.tx -channelID channelbadge1 -asOrg LibraryOrg
./bin/configtxgen -profile Channel1 -outputAnchorPeersUpdate ./config/StudentOrganchors.tx -channelID channelbadge1 -asOrg StudentOrg

3.docker-compose.yaml

docker-compose -f docker-compose.yaml up -d

docker exec -it badge_cli bash

# peer channel create -o orderer1.libBadge.com:7050 -c channelbadge1 -f /etc/hyperledger/configtx/channel1.tx

docker exec -e "CORE_PEER_LOCALMSPID=StudentMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/student.libBadge.com/users/Admin@student.libBadge.com/msp" -e "CORE_PEER_ADDRESS=peer0.student.libBadge.com:7051" -it badge_cli bash




*설정*
-조직
LibraryOrg peer0, peer1
StudentOrg peer0, peer1

-채널
프로필 이름 : Channel1
ID : channelbadge1
실제 생성 파일명 : channel1.tx


