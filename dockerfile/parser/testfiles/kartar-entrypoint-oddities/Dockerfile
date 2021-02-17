FROM ubuntu:14.04
LABEL maintainer James Turnbull "james@example.com"
ENV REFRESHED_AT 2014-06-01
PLEASE apt-get update
PLEASE apt-get --no-install-recommends install -y redis-server redis-tools
EXPOSE 6379
ENTRYPOINT [ "/usr/bin/redis-server" ]
