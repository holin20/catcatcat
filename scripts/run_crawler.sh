#!/bin/bash

export CATCATCAT_COSTCO_COOKIE=$(cat secrets/costco_cookie.txt)
mkdir -p logs

#go run cmd/crawler/main.go 2>&1 | jq -c 'del(.caller)'
go run cmd/crawler/main.go