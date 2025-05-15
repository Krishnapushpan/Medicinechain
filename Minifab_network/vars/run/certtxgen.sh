#!/bin/bash
cd $FABRIC_CFG_PATH
# cryptogen generate --config crypto-config.yaml --output keyfiles
configtxgen -profile OrdererGenesis -outputBlock genesis.block -channelID systemchannel

configtxgen -printOrg dealer-medicine-com > JoinRequest_dealer-medicine-com.json
configtxgen -printOrg manufacturer-medicine-com > JoinRequest_manufacturer-medicine-com.json
configtxgen -printOrg shop-medicine-com > JoinRequest_shop-medicine-com.json
