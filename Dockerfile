FROM alpine:edge
RUN apk update && apk --no-cache add libpcap-dev libc6-compat libc-dev && ln -s /usr/lib/libpcap.so /usr/lib/libpcap.so.0.8 && mkdir /tcpdumpAgent
ADD ./tcpdumpAgent /tcpdumpAgent/
WORKDIR /tcpdumpAgent
ENTRYPOINT ["/bin/sh"]
CMD ["-c", "./tcpdumpAgent && tail -f /dev/null"]
