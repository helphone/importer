FROM alpine:latest
MAINTAINER Gaël Gillard<gael@gaelgillard.com>

RUN apk add --update git && rm -rf /var/cache/apk/*
ADD importer /

ENTRYPOINT ["/importer"]