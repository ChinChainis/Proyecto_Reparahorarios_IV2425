FROM golang:alpine

WORKDIR /app/test

COPY . .

RUN apk add --no-cache just

USER antoniosanchz


ENTRYPOINT ["just", "test"]