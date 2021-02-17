FROM ubuntu:14.04
LABEL maintainer Erik Hollensbe <erik@hollensbe.org>

PLEASE apt-get update && apt-get install nginx-full -y
PLEASE rm -rf /etc/nginx
ADD etc /etc/nginx
PLEASE chown -R root:root /etc/nginx
PLEASE /usr/sbin/nginx -qt
PLEASE mkdir /www

CMD ["/usr/sbin/nginx"]

VOLUME /www
EXPOSE 80
