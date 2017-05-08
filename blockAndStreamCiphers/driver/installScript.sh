#!/bin/bash

echo "running blockAndStreamCiphers script"
echo "building functions"
cd ../
go install

echo "building driver"
cd ./driver
go install

echo "build successful, running program..."
cd ../../../../../bin/
./driver
