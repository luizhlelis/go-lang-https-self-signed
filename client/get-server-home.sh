#!/bin/ash
echo "Installing curl package"
apk update && apk add curl && rm -rf /var/cache/apk/*
echo "Two requests below to get https server home"
sleep 10
curl https://https-server:8081/home
sleep 20
curl https://https-server:8081/home