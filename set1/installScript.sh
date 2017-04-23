#!/bin/bash

echo "running set1 script"
echo "building set1 functions"
cd ./functions
go install

echo "building driver"
cd ../driver
go install

echo "build successful, running program..."
cd ../../../../../bin/
./driver
