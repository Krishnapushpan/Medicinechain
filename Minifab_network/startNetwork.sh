#!/bin/sh

echo "Start the Network"
minifab netup -s couchdb -e true -i 2.4.8 -o manufacturer.medicine.com

sleep 5

echo "create the channel"
minifab create -c medicinechannel

sleep 2

echo "Join the channel"
minifab join -c medicinechannel

sleep 2

echo "achor tag update"
minifab anchorupdate