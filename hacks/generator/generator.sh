#!/usr/bin/env bash

terraform fmt -recursive
terraform init
terraform apply --var target=./../../generated/bungie/data -auto-approve
go fmt ./../../generated/bungie/data
