FROM ubuntu

ARG DATE

RUN echo ${DATE} > date.txt

RUN mkdir /scripts
RUN mkdir /scripts/long
WORKDIR /scripts
COPY tests.sh .
COPY output .
COPY long/main ./long

RUN ./long/main

