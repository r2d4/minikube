FROM alpine:3.4

RUN curl -sSL -o /peer-finder https://storage.googleapis.com/kubernetes-release/pets/peer-finder

COPY on-start.sh /
COPY install.sh /

RUN chmod -c 775 /install.sh /on-start.sh /peer-finder

ENTRYPOINT ["/install.sh"]