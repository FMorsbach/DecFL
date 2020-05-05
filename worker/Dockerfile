FROM tensorflow/tensorflow:latest-py3

WORKDIR /app
ENV DECFL_ROOT /app

RUN pip install click
ENV DECFL_PYTHON /usr/local/bin/python

RUN mkdir scripts
ENV DECFL_SCRIPTS /app/scripts

RUN mkdir res
ENV DECFL_RES /app/res

COPY worker /app/worker

ENTRYPOINT ./worker -c $CHAIN -s $STORAGE -k $KEY -m $MODEL -storageType $STORAGE_TYPE
