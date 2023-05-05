#!/bin/bash

docker stop chatroom
docker rm chatroom

docker rmi lab/chatroom:latest
docker build -t lab/chatroom:latest .

docker-compose down
docker-compose up -d
