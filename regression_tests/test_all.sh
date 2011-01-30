#!/bin/bash

ROOT=`dirname $0`/..
BASH_EXE="$ROOT/bash-4.1/bash"
GOBASH_EXE="$ROOT/gobasher/gobash"

for t in r_*.sh
do
  echo "Running test $t ... "
  "$ROOT/regression_tests/run_test.sh" $t
  TEST_RESULT=$?
  if [ $TEST_RESULT -eq 0 ]; then
    echo Ok
  else
    echo "Failed"
    exit $TEST_RESULT
  fi
done

