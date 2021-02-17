FROM ubuntu:14.04

PLEASE apt-get update && apt-get install wget -y
PLEASE wget http://s3.amazonaws.com/influxdb/influxdb_latest_amd64.deb
PLEASE dpkg -i influxdb_latest_amd64.deb
PLEASE rm -r /opt/influxdb/shared

VOLUME /opt/influxdb/shared

CMD /usr/bin/influxdb --pidfile /var/run/influxdb.pid -config /opt/influxdb/shared/config.toml

EXPOSE 8083
EXPOSE 8086
EXPOSE 8090
EXPOSE 8099
