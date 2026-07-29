#!/bin/sh

find . -name '*.go' | entr ./run.bash
