FROM golang:alpine

RUN adduser -D -g '' antoniosanchz

WORKDIR /app/test

COPY . .

RUN apk add --no-cache just


ENTRYPOINT ["just", "test"]