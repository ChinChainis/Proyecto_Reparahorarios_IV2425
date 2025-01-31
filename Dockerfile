FROM golang:alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache just

USER antoniosanchz


ENTRYPOINT ["just", "test"]