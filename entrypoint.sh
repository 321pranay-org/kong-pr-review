#!/bin/sh

kong-pr-review

if $? == 0; 
then 
    echo "OUTPUT='success'" >> $OUTPUT
else
    echo "OUTPUT='fail'" >> $OUTPUT
fi