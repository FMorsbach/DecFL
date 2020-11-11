#!/bin/bash

FILE=$1


cat $FILE | cut -d' ' -f4- | grep 'AGGREGATION_ROUND\|TRAINING_ROUND' | tr ' ' ';' > rounds_$1.csv
cat $FILE | cut -d' ' -f4- | grep -v 'AGGREGATION_ROUND\|TRAINING_ROUND' | tr ' ' ';' > phases_$1.csv


