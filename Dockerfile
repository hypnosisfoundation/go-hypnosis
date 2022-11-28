# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.16-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

ADD . /go-hypnosis
RUN cd /go-hypnosis && go run build/ci.go install ./cmd/geth
RUN cp build/bin/geth build/bin/hypnosis
RUN rm build/bin/geth

# Pull Hypnosis into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-hypnosis/build/bin/hypnosis /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["hypnosis"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"
