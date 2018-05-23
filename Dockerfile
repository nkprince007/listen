FROM alpine:3.2
ADD listen /listen
WORKDIR /
ENTRYPOINT [ "/listen" ]
