FROM golang:alpine as build-env
LABEL MAINTAINER Naveen Kumar Sangi <naveenkumarsangi@protonmail.com>

RUN apk --no-cache add git
ENV SOURCE /go/src/gitlab.com/nkprince007/listen
RUN mkdir -p $SOURCE
ADD . $SOURCE
WORKDIR $SOURCE
RUN go get -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest

RUN apk --no-cache add ca-certificates
ENV SOURCE /go/src/gitlab.com/nkprince007/listen
WORKDIR /root/
COPY --from=build-env $SOURCE .
CMD ["./app"]
