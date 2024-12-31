#!/bin/bash

export CATCATCAT_COSTCO_COOKIE=$(cat secrets/costco_cookie.txt)
go run cmd/crawler/main.go 2>&1 | jq -c 'del(.caller)'