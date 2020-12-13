#!/bin/ash
#   Installing ca-certificates package which comes with preinstalled certs and enables update with your third party certs
echo "Installing ca-certificates package"
apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
echo "Waiting for the server to server create the certificate"
sleep 20
echo "Copy certificate to ca-certificates folder"
cp ../certificates/servercert.crt /usr/local/share/ca-certificates/servercert.crt
#   The command below will make the client container to trust in server cert
echo "Trusting server self signed cert"
update-ca-certificates