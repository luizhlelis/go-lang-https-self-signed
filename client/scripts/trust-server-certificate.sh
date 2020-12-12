#!/bin/ash
#   Installing ca-certificates package which comes with preinstalled certs and enables update with your third party certs
echo "Installing ca-certificates package"
apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
#   The server cert is already in the folder /usr/local/share/ca-certificates/ (as cofigured in volume).
#   The command below will make the client container to trust in server cert
echo "Trusting server self signed cert"
sudo update-ca-certificates