FROM tensorflow/tensorflow:2.3.1

# Preload datasets
RUN python3 -c "import tensorflow as tf; tf.keras.datasets.cifar10.load_data()"
RUN python3 -c "import tensorflow as tf; tf.keras.datasets.mnist.load_data()"

WORKDIR /app
ENV DECFL_ROOT /app

RUN pip install click
ENV DECFL_PYTHON /usr/local/bin/python

RUN mkdir scripts
ENV DECFL_SCRIPTS /app/scripts

RUN mkdir res
ENV DECFL_RES /app/res

COPY worker /app/worker

CMD /app/worker -c ${CHAIN} -s ${STORAGE} -k ${KEY} -m ${MODEL} -storageType ${STORAGE_TYPE}
