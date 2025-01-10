FROM docker.elastic.co/beats/filebeat:8.10.0

COPY filebeat.yml /usr/share/filebeat/filebeat.yml

USER root

RUN chmod 644 /usr/share/filebeat/filebeat.yml

USER root
