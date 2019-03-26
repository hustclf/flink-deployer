FROM golang:1.10.4-alpine3.8 as build
WORKDIR /go/src/flink-deployer/
COPY . .
RUN go build ./cmd/cli

FROM alpine:3.8
WORKDIR /flink-deployer
COPY --from=build /go/src/flink-deployer/cli .
VOLUME [ "/data/flink" ]
#ENTRYPOINT [ "/flink-deployer/cli" ]
CMD [ "help" ]
