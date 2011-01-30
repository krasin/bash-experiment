#!/bin/bash

ROOT=`dirname $0`/..
BASH_EXE="$ROOT/bash-4.1/bash"
GOBASH_EXE="$ROOT/gobasher/gobash"

echo "Root: " $ROOT
echo "Bash executable: " $BASH_EXE
echo "Gobash executable: " $GOBASH_EXE

TEST_SCRIPT="$1"
BASH_STDERR="$(mktemp)"
echo "BASH_STDERR: " $BASH_STDERR
GOBASH_STDERR="$(mktemp)"
echo "GOBASH_STDERR: " $GOBASH_STDERR

echo "Running bash.."
$BASH_EXE $TEST_SCRIPT >/dev/null 2> $BASH_STDERR
BASH_RESULT=$?
if [ ! $BASH_RESULT -eq 0 ]; then
  return $BASH_RESULT
fi
echo "Running gobash..."
$GOBASH_EXE $TEST_SCRIPT >/dev/null 2> $GOBASH_STDERR
GOBASH_RESULT=$?
if [ ! $GOBASH_RESULT -eq 0 ]; then
  return $GOBASH_RESULT
fi
diff $BASH_STDERR $GOBASH_STDERR
DIFF_RESULT=$?
echo diff returned $DIFF_RESULT
exit $DIFF_RESULT

