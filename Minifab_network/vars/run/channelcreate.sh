#!/bin/bash
# Script to create channel block 0 and then create channel
cp $FABRIC_CFG_PATH/core.yaml /vars/core.yaml
cd /vars
export FABRIC_CFG_PATH=/vars
configtxgen -profile OrgChannel \
  -outputCreateChannelTx medicinechannel.tx -channelID medicinechannel

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ID=cli
export CORE_PEER_ADDRESS=192.168.43.197:7003
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/manufacturer.medicine.com/peers/peer1.manufacturer.medicine.com/tls/ca.crt
export CORE_PEER_LOCALMSPID=manufacturer-medicine-com
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/manufacturer.medicine.com/users/Admin@manufacturer.medicine.com/msp
export ORDERER_ADDRESS=192.168.43.197:7011
export ORDERER_TLS_CA=/vars/keyfiles/ordererOrganizations/medicine.com/orderers/orderer3.medicine.com/tls/ca.crt
peer channel create -c medicinechannel -f medicinechannel.tx -o $ORDERER_ADDRESS \
  --cafile $ORDERER_TLS_CA --tls
