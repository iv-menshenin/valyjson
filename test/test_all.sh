#!/bin/bash

echo "GENERATE test code"
go test -v --run=Test_fld_FillStatements .

echo "TEST..."
cd ./test/tests/fill/
go test -v ./...