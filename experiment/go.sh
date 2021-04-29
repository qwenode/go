#!/bin/sh
cd ../src

sh make.bash

cd ../experiment

../bin/go run main.go