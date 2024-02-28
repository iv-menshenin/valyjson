#!/bin/bash

ROOT=$(pwd)

echo "##teamcity[testSuiteStarted name='generator.UnitTests']"
cd $ROOT/generator || exit 1
go test -json -v ./... || echo "##teamcity[testFailed name='generator.UnitTests']"
echo "##teamcity[testSuiteFinished name='generator.UnitTests']"

echo "##teamcity[blockOpened name='Tests_Preparation']"
cd $ROOT/test || exit 1
go test -json -v --run=Test_fld_FillStatements . || echo "##teamcity[testFailed name='Tests_Preparation_1']"
go test -json -v --run=Test_GenerateVJson . || echo "##teamcity[testFailed name='Tests_Preparation_2']"
echo "##teamcity[blockClosed name='Tests_Preparation']"

echo "##teamcity[blockOpened name='Tests_Execution']"

echo "##teamcity[testSuiteStarted name='generator.GeneratedCode']"
cd $ROOT/test/tests/fill/ || exit 1
go test -json -v ./... || echo "##teamcity[testFailed name='generator.GeneratedCode']"
echo "##teamcity[testSuiteFinished name='generator.GeneratedCode']"

echo "##teamcity[testSuiteStarted name='generator.Functional']"
cd $ROOT/test/vjson/ || exit 1
go test -race -json -v ./... || echo "##teamcity[testFailed name='generator.Functional']"
echo "##teamcity[testSuiteFinished name='generator.Functional']"

echo "##teamcity[testSuiteStarted name='generator.RaceTests']"
cd $ROOT/test/race/ || exit 1
go test -race -json -v ./... || echo "##teamcity[testFailed name='generator.RaceTests']"
echo "##teamcity[testSuiteFinished name='generator.RaceTests']"

echo "##teamcity[blockClosed name='Tests_Execution']"
