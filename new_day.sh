#!/usr/bin/env bash

if [ -z $1 ]; then exit 1; fi

year=2023
day=$1

if [ ! -z $2 ]; then year=$1; day=$2; fi

echo "Setting up day $day of $year"

touch ./inputs/$year/day$day{-test,}.txt
cp -n ./$year/dayXX.go.template ./$year/day$day.go
sed -i "s/XX/$day/g" ./$year/day$day.go
