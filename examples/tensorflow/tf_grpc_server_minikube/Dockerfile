FROM tensorflow/tensorflow:0.12.0

COPY tf_grpc.py /

RUN chmod -c 775 /tf_grpc.py

ENTRYPOINT ["/tf_grpc.py"]

