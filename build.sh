#!/bin/bash

chmod +x create.sh
sh create.sh

BASE_PATH=$(pwd)

make build

chmod +x docker.sh
sh docker.sh
