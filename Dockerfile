FROM alpine:3.2
MAINTAINER Gaël Gillard<gael@gaelgillard.com>

RUN apk add --update git && \
		git clone https://github.com/helphone/data.git /etc/data && \
		rm -rf /var/cache/apk/*
ADD importer /bin/

CMD ["/bin/importer"]
