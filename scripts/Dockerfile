FROM alpine:3.3
MAINTAINER Gaël Gillard<gael@gaelgillard.com>

RUN apk add --update git && \
		rm -rf /var/cache/apk/*

COPY importer /bin/

CMD ["/bin/importer"]
