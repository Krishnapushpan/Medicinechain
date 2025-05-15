#!/bin/bash
# Script to join a peer to a channel
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_ID=cli
export CORE_PEER_ADDRESS=192.168.43.197:7003
export CORE_PEER_TLS_ROOTCERT_FILE=/vars/keyfiles/peerOrganizations/manufacturer.medicine.com/peers/peer1.manufacturer.medicine.com/tls/ca.crt
export CORE_PEER_LOCALMSPID=manufacturer-medicine-com
export CORE_PEER_MSPCONFIGPATH=/vars/keyfiles/peerOrganizations/manufacturer.medicine.com/users/Admin@manufacturer.medicine.com/msp
export ORDERER_ADDRESS=192.168.43.197:7011
export ORDERER_TLS_CA=/vars/keyfiles/ordererOrganizations/medicine.com/orderers/orderer3.medicine.com/tls/ca.crt
if [ ! -f "medicinechannel.genesis.block" ]; then
  peer channel fetch oldest -o $ORDERER_ADDRESS --cafile $ORDERER_TLS_CA \
  --tls -c medicinechannel /vars/medicinechannel.genesis.block
fi

peer channel join -b /vars/medicinechannel.genesis.block \
  -o $ORDERER_ADDRESS --cafile $ORDERER_TLS_CA --tls
