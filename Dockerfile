FROM golang:1.15.5-alpine3.12

WORKDIR /app

RUN sudo apt update && sudo apt install -y redis-server
