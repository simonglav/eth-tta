# syntax=docker/dockerfile:1
FROM golang:1.17-alpine
WORKDIR /eth-tta
COPY . .
RUN go mod download
RUN go build ./cmd/apiserver
CMD ["./apiserver"]