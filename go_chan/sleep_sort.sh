#!/bin/bash

for n in $@; do
    (sleep $n && echo $n)&
done
wait