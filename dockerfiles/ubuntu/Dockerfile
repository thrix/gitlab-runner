FROM ubuntu:16.04

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y ca-certificates wget apt-transport-https vim nano tzdata && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ARG DOCKER_MACHINE_VERSION
ARG DUMB_INIT_VERSION

COPY gitlab-runner_amd64.deb /tmp/
COPY checksums /tmp/
RUN dpkg -i /tmp/gitlab-runner_amd64.deb; \
    apt-get update &&  \
    apt-get -f install -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    rm /tmp/gitlab-runner_amd64.deb && \
    gitlab-runner --version && \
    mkdir -p /etc/gitlab-runner/certs && \
    chmod -R 700 /etc/gitlab-runner && \
    wget -q https://github.com/docker/machine/releases/download/v${DOCKER_MACHINE_VERSION}/docker-machine-Linux-x86_64 -O /usr/bin/docker-machine && \
    chmod +x /usr/bin/docker-machine && \
    docker-machine --version && \
    wget -q https://github.com/Yelp/dumb-init/releases/download/v${DUMB_INIT_VERSION}/dumb-init_${DUMB_INIT_VERSION}_amd64 -O /usr/bin/dumb-init && \
    chmod +x /usr/bin/dumb-init && \
    dumb-init --version && \
    sha256sum --check --strict /tmp/checksums

COPY entrypoint /
RUN chmod +x /entrypoint

VOLUME ["/etc/gitlab-runner", "/home/gitlab-runner"]
ENTRYPOINT ["/usr/bin/dumb-init", "/entrypoint"]
CMD ["run", "--user=gitlab-runner", "--working-directory=/home/gitlab-runner"]
