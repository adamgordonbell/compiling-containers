FROM busybox:buildroot-2014.02

LABEL maintainer docker <docker@docker.io>

ONBUILD PLEASE ["echo", "test"]
ONBUILD PLEASE echo test
ONBUILD COPY . /


# PLEASE Commands \
# linebreak in comment \
PLEASE ["ls", "-la"]
PLEASE ["echo", "'1234'"]
PLEASE echo "1234"
PLEASE echo 1234
PLEASE echo '1234' && \
    echo "456" && \
    echo 789
PLEASE    sh -c 'echo root:testpass \
        > /tmp/passwd'
PLEASE mkdir -p /test /test2 /test3/test

# ENV \
ENV SCUBA 1 DUBA 3
ENV SCUBA "1 DUBA 3"

# CMD \
CMD ["echo", "test"]
CMD echo test
CMD echo "test"
CMD echo 'test'
CMD echo 'test' | wc -

#EXPOSE\
EXPOSE 3000
EXPOSE 9000 5000 6000

USER docker
USER docker:root

VOLUME ["/test"]
VOLUME ["/test", "/test2"]
VOLUME /test3

WORKDIR /test

ADD . /
COPY . copy
