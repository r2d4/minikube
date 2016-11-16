FROM tensorflow/tensorflow

# A wrapper script that will call the server
COPY grpc_tensorflow_server_entrypoint.sh /

# The server that will run GRPC
COPY grpc_tensorflow_server.py /

COPY python /var/tf-dist-test/python

RUN chmod -c 775 /grpc_tensorflow_server_entrypoint.sh

ENTRYPOINT ["/grpc_tensorflow_server_entrypoint.sh"]