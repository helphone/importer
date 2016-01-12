FROM alpine:3.3
MAINTAINER Gaël Gillard<gael@gaelgillard.com>

RUN echo http://nl.alpinelinux.org/alpine/edge/testing >> /etc/apk/repositories && \
		apk add --update libgit2 && \
		rm -rf /var/cache/apk/*

ADD importer /bin/

CMD ["/bin/importer"]
