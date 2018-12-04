FROM scratch
ADD echo-server /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/echo-server"]