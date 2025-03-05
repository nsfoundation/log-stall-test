FROM ubuntu

ARG DATE

RUN echo ${DATE} > date.txt

RUN mkdir /scripts
WORKDIR /scripts
COPY tests.sh .
COPY output .
RUN ./tests.sh 

