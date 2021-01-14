#!/bin/bash
# Install test-reporter
curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter
chmod +x ./cc-test-reporter

# before build step
./cc-test-reporter before-build

# RUN TEST HERE
go test ./... -coverprofile c.out

# after build step
./cc-test-reporter after-build
