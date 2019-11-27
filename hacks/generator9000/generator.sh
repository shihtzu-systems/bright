#!/usr/bin/env bash

terraform fmt -recursive
terraform init
terraform apply --var target=./../generator -auto-approve -auto-approve
terraform fmt ./../generator