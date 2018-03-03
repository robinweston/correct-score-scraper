#!/bin/sh 
set -e

docker run  --rm -it -v "$(pwd)":/go/src/app correct-score-image $@