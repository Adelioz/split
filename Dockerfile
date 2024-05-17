# syntax=docker/dockerfile:1

FROM golang:1.21.10-alpine as go-base

WORKDIR /build
COPY . .

RUN go build -o /split ./cmd/split

FROM alpine:latest

WORKDIR /app
COPY --from=go-base /split /app/

CMD [ "/app/split", "-devel" ]