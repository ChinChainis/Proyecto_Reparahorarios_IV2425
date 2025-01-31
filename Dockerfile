FROM golang:alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache just

USER tester


ENTRYPOINT ["just", "test"]