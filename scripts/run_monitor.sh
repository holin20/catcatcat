#!/bin/bash

export CATCATCAT_GAPP_PW=$(cat secrets/google_app_password.txt)
mkdir -p logs

go run cmd/monitor/main.go 2>&1 | jq -c 'del(.caller)'
#go run cmd/monitor/main.go