#!/bin/bash

DIR=$1
TMP_FILE1="/tmp/decfl_tmp1"
TMP_FILE2="/tmp/decfl_tmp2"
TMP_FILE3="/tmp/decfl_tmp3"
TMP_FILE4="/tmp/decfl_tmp4"
REGEX='AGGREGATION_ROUND\|TRAINING_ROUND\|WAITING_AFTER_'


echo "" > $TMP_FILE2
echo "" > $TMP_FILE4

for f in $DIR/*; do

    cat $f | cut -d' ' -f4- | grep $REGEX | tr ' ' ';' > $TMP_FILE1
    cat $f | cut -d' ' -f4- | grep -v $REGEX | tr ' ' ';' > $TMP_FILE3

    EXP_ID=$(echo $f | cut -d'/' -f2 | tr '_' ';' | cut -d'.' -f1)

    sed -e "s/$/;$EXP_ID/" -i $TMP_FILE1
    sed -e "s/$/;$EXP_ID/" -i $TMP_FILE3

    cat $TMP_FILE1 >> $TMP_FILE2
    cat $TMP_FILE3 >> $TMP_FILE4
done


cat $TMP_FILE2 | (echo -n "STEP;ROUND;TIME;MODEL;CLIENTS;NODE_ID" && cat ) > rounds.csv
cat $TMP_FILE4 | (echo -n "STEP;TIME;MODEL;CLIENTS;NODE_ID" && cat ) > phases.csv

