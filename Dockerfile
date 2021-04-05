FROM alpine:edge
RUN apk update && apk --no-cache add tcpdump jq docker  && mkdir /containerTcpdump
ADD ./containerTcpdump /containerTcpdump/
WORKDIR /containerTcpdump
ENTRYPOINT ["/bin/sh"]
CMD ["-c", "./containerTcpdump && tail -f /dev/null"]
