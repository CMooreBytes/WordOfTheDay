
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/github.com/cmoorebytes/wordoftheday
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/wordoftheday /app
COPY ./wwwroot /wwwroot
ENTRYPOINT ./app
LABEL Name=go Version=0.0.1
EXPOSE 8000

#docker build -t wordoftheday .
#docker run --publish 8000:5000 --name test --rm wordoftheday