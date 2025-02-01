FROM golang:alpine

WORKDIR /app/test

COPY . .

RUN apk add --no-cache just


ENTRYPOINT ["just", "test"]