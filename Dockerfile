FROM golang:1.17.3-alpine3.15 AS build-env

# Set up dependenciesbuild-env
ENV PACKAGES git build-base

# Set working directory for the build
WORKDIR /go/src/github.com/bdengine/bde

RUN echo -e http://mirrors.tuna.tsinghua.edu.cn/alpine/v3.11/main/ > /etc/apk/repositories
# Install dependencies
RUN apk add --update $PACKAGES
RUN apk add linux-headers

# Add source files
COPY . .
ENV GOPROXY=https://goproxy.io
RUN go env
# Make the binary
RUN mkdir build
RUN go build  -o ./build/ ./cmd/bdeosd

# Final image
FROM alpine

RUN echo -e http://mirrors.tuna.tsinghua.edu.cn/alpine/v3.11/main/ > /etc/apk/repositories
# Install ca-certificates
RUN apk add --update ca-certificates jq
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/src/github.com/bdengine/bde/build/bdeosd /usr/bin/bdeosd
#COPY --from=build-env /go/src/github.com/tharsis/evmos/scripts/dockerStart.sh /var/evmosd/dockerStart.sh

#RUN chmod +x /var/evmosd/dockerStart.sh
ENV CHAINID=bdeos_18888-1
ENV MONIKER=sc-net
ENV TOKENNAME=aphoton
ENV KEY=shangchain
ENV KEYRING=test
ENV NODEHOME=/root/.bdeosd/
EXPOSE 26656 26657 1317 9090

