FROM golang:1.23.0

COPY . /app

WORKDIR /app

RUN ["/bin/bash"]