FROM tensorflow/tensorflow

# A wrapper script that will call the server
COPY grpc_tensorflow_server_wrapper.sh

# The server that will run GRPC
COPY grpc_tensorflow_server.py


ENTRYPOINT 