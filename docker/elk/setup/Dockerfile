# ARG ELASTIC_VERSION
ARG ELASTIC_VERSION=8.5.3
# https://www.docker.elastic.co/
FROM elasticsearch:${ELASTIC_VERSION}

USER root

RUN set -eux; \
	mkdir /state; \
	chmod 0775 /state; \
	chown elasticsearch:root /state

USER elasticsearch:root

ENTRYPOINT ["/entrypoint.sh"]
