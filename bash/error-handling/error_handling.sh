#!/usr/bin/env bash


if [ $# -gt 0 ]
then
    if [ $# -gt 1 ]
    then
        echo "Too many arguments"
        exit 1
    fi
    echo "Hello, $1"
else
    echo "Usage: ./error_handling <greetee>"
    exit 1
fi