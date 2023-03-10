#!/bin/bash

echo "##teamcity[blockOpened name='Tests_Preparation']"
go test -json -v --run=Test_fld_FillStatements . || echo "##teamcity[testFailed name='Tests_Preparation']"
echo "##teamcity[blockClosed name='Tests_Preparation']"

echo "##teamcity[blockOpened name='Tests_Execution']"

echo "##teamcity[testSuiteStarted name='generator.GeneratedCode']"
cd ./test/tests/fill/ || exit 1
go test -json -v ./... || echo "##teamcity[testFailed name='generator.GeneratedCode']"
echo "##teamcity[testSuiteFinished name='generator.GeneratedCode']"

echo "##teamcity[testSuiteStarted name='generator.UnitTests']"
cd ../../../generator || exit 1
go test -json -v ./... || echo "##teamcity[testFailed name='generator.UnitTests']"
echo "##teamcity[testSuiteFinished name='generator.UnitTests']"

echo "##teamcity[blockClosed name='Tests_Execution']"
