REPO_NAME=harbor
HARBOR_VERSION=0.3.1
ALPINE_VERSION=3.6

all: test

check:
	curl -s "http://registry.docker.elo7aws.com.br/v2/${REPO_NAME}/tags/list" \
		| jq -e --arg tag ${HARBOR_VERSION}-alpine${ALPINE_VERSION} '.tags[] | select(.==$$tag)' &> /dev/null && exit 1 || exit 0

test:
	harbor -e HARBOR_VERSION=${HARBOR_VERSION} -e ALPINE_VERSION=${ALPINE_VERSION} -e REPO_NAME=${REPO_NAME} --no-docker-push --no-latest-tag

prod: check
	harbor -e HARBOR_VERSION=${HARBOR_VERSION} -e ALPINE_VERSION=${ALPINE_VERSION} -e REPO_NAME=${REPO_NAME} --no-latest-tag

force-prod:
	harbor -e HARBOR_VERSION=${HARBOR_VERSION} -e ALPINE_VERSION=${ALPINE_VERSION} -e REPO_NAME=${REPO_NAME} --no-latest-tag