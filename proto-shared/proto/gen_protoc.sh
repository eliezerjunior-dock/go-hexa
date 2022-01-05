#!/bin/bash

mkdir -p ../generated/go
protoc rate/rate.proto --go_out=plugins=grpc:../../../../
protoc balance/balance.proto --go_out=plugins=grpc:../../../../