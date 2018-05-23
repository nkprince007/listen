FROM golang:onbuild
LABEL MAINTAINER Naveen Kumar Sangi <naveenkumarsangi@protonmail.com>
ENV SOURCE /go/src/gitlab.com/nkprince007/listen/
ADD . $SOURCE
RUN go get -u -v .
WORKDIR $SOURCE
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/gitlab.com/nkprince007/listen/app .
CMD ["./app"]
