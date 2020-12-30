FROM ubuntu:22.04

ARG PENGUIN_VERSION=1.0.1

RUN mkdir -p /app/configs
RUN mkdir -p /app/var/logs
RUN apt-get update

WORKDIR /app

RUN curl -sL https://github.com/uptimedog/penguin/releases/download/v${PENGUIN_VERSION}/penguin_${PENGUIN_VERSION}_Linux_x86_64.tar.gz | tar xz
RUN rm LICENSE
RUN rm README.md

COPY ./config.dist.yml /app/configs/

EXPOSE 8000

VOLUME /app/configs
VOLUME /app/var

RUN ./penguin version

CMD ["./penguin", "run", "-c", "/app/configs/config.dist.yml"]