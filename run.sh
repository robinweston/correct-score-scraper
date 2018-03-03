#!/bin/sh 
set -e

docker build -t correct-score-image .
docker run  --rm -it correct-score-image