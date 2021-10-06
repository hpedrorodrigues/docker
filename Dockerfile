FROM golang:1.17.1-alpine3.14

COPY . /app
WORKDIR /app

ENTRYPOINT ["/app/run.sh"]
