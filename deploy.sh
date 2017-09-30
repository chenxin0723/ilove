#!/usr/bin/env bash

if go run main.go -compile-templates=true ; then
    # harp -s $env kill

    echo "Deploying----------------------- prod site"
    harp -s prod deploy
     harp -s prod log
    # echo "Deploying----------------------- draft site"
    # harp -s draft deploy
    #  harp -s draft log

else
    echo "Failed: go run main.go -compile-templates=true"
    echo "Please try: source config/.envrc"
fi
