#!/bin/bash

echo "running set2 script"
echo "building set2 functions"
cd ./functions
go install

echo "building driver"
cd ../driver
go install

echo "build successful, running program..."
cd ../../../../../bin/
./driver
