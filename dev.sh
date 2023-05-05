#!/bin/bash

chmod +x create.sh
sh create.sh

BASE_PATH=$(pwd)

make build-dev

chmod +x docker.sh
sh docker.sh
