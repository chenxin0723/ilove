#!/usr/bin/env bash

env=dev


if go run main.go -compile-templates=true ; then
    # harp -s $env kill
    harp -s $env deploy
    harp -s $env log
else
    echo "Failed: go run main.go -compile-templates=true"
    echo "Please try: source config/.envrc"
fi
