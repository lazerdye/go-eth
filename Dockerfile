FROM golang:1.17-bullseye AS builder

RUN mkdir /tjoe-go

COPY . /tjoe-go

RUN cd /tjoe-go && go build -o /tjoe-go/bin/eth-cli cli/*.go

FROM debian:bullseye-slim

COPY --from=builder /tjoe-go/bin/eth-cli /usr/local/bin/eth-cli

ENTRYPOINT ["/usr/local/bin/eth-cli"]
