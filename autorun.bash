#!/bin/sh

while true; do
    find . -name '*.go' | entr -d ./run.bash
    if [ "$?" = 2 ]; then
        continue
    fi
done
