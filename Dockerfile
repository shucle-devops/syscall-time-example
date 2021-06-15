#build stage
FROM golang:alpine
RUN apk add --no-cache git
WORKDIR /go/src/this-example
COPY . .
RUN go get -d -v ./...