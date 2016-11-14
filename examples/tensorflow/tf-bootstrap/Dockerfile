FROM gcr.io/google_containers/peer-finder:0.1

COPY on-start.sh /

RUN chmod -c 775 /on-start.sh /peer-finder

ENTRYPOINT ["/peer-finder"]