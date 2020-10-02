#!/bin/bash

coverage() {
 go test ./... -coverprofile cover.out -coverpkg "github.com/chaitya62/noobdb/type" "github.com/chaitya62/noobdb/tests/type"  >> /dev/null 2>&1
 go tool cover -func cover.out
 rm cover.out
}

run_test() {
  go test ./... -v
}

unset ACTION

while getopts 'ct' c
do
  case $c in
    c) ACTION=COVERAGE;;
    t) ACTION=RUNTEST
  esac
done

# this only does one thing at a time 
case $ACTION in
  COVERAGE) coverage;;
  RUNTEST) run_test
esac
