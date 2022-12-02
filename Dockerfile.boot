FROM ubuntu:18.04

RUN cd /home/
RUN git clone https://hypnosisfoundation/go-hypnosis
RUN cd go-hypnosis
RUN make hypnosis

RUN 
