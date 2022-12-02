FROM ubuntu:18.04

ARG HYPNOSIS_PASS
ARG HYPNOSIS_PUBKEY

RUN cd /home/
RUN sudo apt-get install git golang make
RUN git clone https://hypnosisfoundation/go-hypnosis
RUN cd go-hypnosis
RUN make hypnosis

RUN mkdir -p build/bin/data
RUN ./build/bin/hypnosis --datadir data/ account new

