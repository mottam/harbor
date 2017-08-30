ARG ALPINE_VERSION
FROM alpine:${ALPINE_VERSION}
MAINTAINER BeyondTheWall <engenharia.beyondthewall@elo7.com>

ARG HARBOR_VERSION
ENV HARBOR_VERSION=${HARBOR_VERSION}

ENV DOCKER_VERSION 1.12.3
ENV DOCKER_CERT_PATH /etc/docker/ssl
ENV DOCKER_HOST tcp://daemon.docker.elo7aws.com.br:2376
ENV DOCKER_TLS_VERIFY 1

# Installing Docker Daemon authentication keys
ADD keys/ca.pem   ${DOCKER_CERT_PATH}/ca.pem
ADD keys/cert.pem ${DOCKER_CERT_PATH}/cert.pem
ADD keys/key.pem  ${DOCKER_CERT_PATH}/key.pem

RUN apk update && apk add docker make bash git openssh-client && rm -rf /var/cache/apk/*

ADD https://github.com/elo7/harbor/releases/download/v${HARBOR_VERSION}/harbor_${HARBOR_VERSION}_linux_amd64.tar.gz /usr/local/bin/

ENTRYPOINT ["harbor"]

