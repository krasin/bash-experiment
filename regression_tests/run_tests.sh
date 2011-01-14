#!/bin/bash

ROOT=`dirname $0`/..
BASH_EXE="$ROOT/bash-4.1/bash"
GOBASH_EXE="$ROOT/gobash/gobash"

echo "Root: " $ROOT
echo "Bash executable: " $BASH_EXE
echo "Gobash executable: " $GOBASH_EXE

function run_regression_test() {
  echo -n "Running test $1 ... "
}

for t in r_*.sh
do
  run_regression_test $t
  TEST_RESULT=$?
  if [ "$TEST_RESULT"="0" ]; then
    echo Ok
  else
    echo "Failed"
    exit $TEST_RESULT
  fi
done

