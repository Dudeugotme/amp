FROM appcelerator/gotools:1.4.0
RUN apk --no-cache add sudo iproute2

# Add Docker client from:
# https://github.com/docker-library/docker/blob/master/1.13/Dockerfile
ENV DOCKER_BUCKET get.docker.com
ENV DOCKER_VERSION 1.13.0
ENV DOCKER_SHA256 fc194bb95640b1396283e5b23b5ff9d1b69a5e418b5b3d774f303a7642162ad6

RUN set -x \
	&& curl -fSL "https://${DOCKER_BUCKET}/builds/Linux/x86_64/docker-${DOCKER_VERSION}.tgz" -o docker.tgz \
	&& echo "${DOCKER_SHA256} *docker.tgz" | sha256sum -c - \
	&& tar -xzvf docker.tgz \
	&& mv docker/* /usr/local/bin/ \
	&& rmdir docker \
	&& rm docker.tgz \
	&& docker -v

COPY docker-entrypoint.sh /usr/local/bin/
COPY sudoers /etc/sudoers.d/amp
RUN chmod 0440 /etc/sudoers.d/amp

RUN adduser -S -s /bin/sh sudoer

# pass commands through docker-entrypoint first for special handling
# it's fine to override entrypoint if not running a docker command
ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["sh"]

