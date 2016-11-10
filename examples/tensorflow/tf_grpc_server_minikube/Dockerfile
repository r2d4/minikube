FROM tensorflow/tf_grpc_test_server

# A wrapper script that will call the server
COPY grpc_tensorflow_server_entrypoint.sh /
RUN chmod -c 775 /grpc_tensorflow_server_entrypoint.sh

ENTRYPOINT ["/grpc_tensorflow_server_entrypoint.sh"]

