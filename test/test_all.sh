#!/bin/bash

echo "##teamcity[blockOpened name='Tests_Preparation']"
go test -v --run=Test_fld_FillStatements .
echo "##teamcity[blockClosed name='Tests_Preparation']"

echo "##teamcity[blockOpened name='Tests_Execution']"

echo "##teamcity[testSuiteStarted name='Execute_Test_Generation']"
cd ./test/tests/fill/
go test -v ./...
echo "##teamcity[testSuiteFinished name='Execute_Test_Generation']"

echo "##teamcity[testSuiteStarted name='Execute_Unit_Tests']"
cd ../../../generator
go test -v ./...
echo "##teamcity[testSuiteFinished name='Execute_Unit_Tests']"

echo "##teamcity[blockClosed name='Tests_Execution']"
