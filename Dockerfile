FROM alpine:edge
RUN apk update && apk --no-cache add libpcap-dev libc6-compat libc-dev && ln -s /usr/lib/libpcap.so /usr/lib/libpcap.so.0.8 && mkdir /dumpAgent
ADD ./dumpAgent /dumpAgent/
WORKDIR /dumpAgent
ENTRYPOINT ["/bin/sh"]
CMD ["-c", "./dumpAgent && tail -f /dev/null"]
