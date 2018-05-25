# Stage - Build
FROM golang:alpine as build-env

LABEL MAINTAINER Naveen Kumar Sangi <naveenkumarsangi@protonmail.com>

RUN apk --no-cache add git
RUN mkdir -p /go/src/gitlab.com/gitmate-micro/listen
ADD . /go/src/gitlab.com/gitmate-micro/listen
WORKDIR /go/src/gitlab.com/gitmate-micro/listen

RUN go get -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -s' -o listen .


# Stage - Deploy
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-env /go/src/gitlab.com/gitmate-micro/listen .

CMD ["./listen"]
