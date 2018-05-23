FROM golang:alpine as build-env
LABEL MAINTAINER Naveen Kumar Sangi <naveenkumarsangi@protonmail.com>

RUN apk --no-cache add curl git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
ENV SOURCE /go/src/app
ADD . $SOURCE
WORKDIR $SOURCE
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-env /go/src/app .
CMD ["./app"]
