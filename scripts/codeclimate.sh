#!/bin/bash
REPORTER=cc-test-reporter
# Install test-reporter if it doesnt exist
if [[ ! -f "$REPORTER"  ]]; then
  curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter
  chmod +x ./cc-test-reporter
fi

# before build step
./cc-test-reporter before-build

# RUN TEST HERE
go test ./... -coverprofile c.out

# after build step
./cc-test-reporter after-build
