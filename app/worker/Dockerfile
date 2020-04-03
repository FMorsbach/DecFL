FROM python:3.6-buster

WORKDIR /app

RUN pip install virtualenv
RUN virtualenv venv
RUN venv/bin/pip install tensorflow click
ENV DECFL_PYTHON /app/venv/bin/python

RUN mkdir scripts
ENV DECFL_SCRIPTS /app/scripts

RUN mkdir res
ENV DECFL_RES /app/res

COPY worker /app/worker

#TODO remove
COPY scripts /app/scripts

CMD [ "/bin/bash" ]


