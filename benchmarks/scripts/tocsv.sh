#!/bin/bash

DIR=$1
TMP_FILE="/tmp/decfl_tmp"
TMP_FILE2="/tmp/decfl_tmp2"
REGEX='AGGREGATION_ROUND\|TRAINING_ROUND\|WAITING_AFTER_'


echo "" > $TMP_FILE2


for f in $DIR/*; do

    cat $f | cut -d' ' -f4- | grep $REGEX | tr ' ' ';' > $TMP_FILE
    #cat $FILE | cut -d' ' -f4- | grep -v $REGEX | tr ' ' ';' > phases_$1.csv

    EXP_ID=$(echo $f | cut -d'/' -f2 | tr '_' ';' | cut -d'.' -f1)
    sed -e "s/$/;$EXP_ID/" -i $TMP_FILE
    #sed -e "s/$/;$EXP_ID/" -i phases_$1.csv

    cat $TMP_FILE >> $TMP_FILE2
done


cat $TMP_FILE2 | (echo "STEP;ROUND;TIME;MODEL;CLIENTS;NODE_ID" && cat ) > rounds.csv
