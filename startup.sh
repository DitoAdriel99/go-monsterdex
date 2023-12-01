#!/usr/bin/env sh

./wait-for-it.sh database:5432 --timeout=50
./app
