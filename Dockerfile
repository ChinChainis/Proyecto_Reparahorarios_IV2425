FROM golang:alpine

RUN useradd -ms /bin/bash antoniosanchz

WORKDIR /app/test

COPY . .

RUN apk add --no-cache just


ENTRYPOINT ["just", "test"]