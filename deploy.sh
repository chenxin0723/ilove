#!/usr/bin/env bash

if go run main.go -compile-templates=true ; then
    # harp -s $env kill

    echo "Deploying----------------------- prod site"
    harp -s prod deploy
    echo "Deploying----------------------- draft site"
    harp -s draft deploy
    harp -s prod log
else
    echo "Failed: go run main.go -compile-templates=true"
    echo "Please try: source config/.envrc"
fi
