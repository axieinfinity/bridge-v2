FROM golang:1.18-buster as builder

WORKDIR /opt

ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

COPY . /opt/bridge
RUN cd bridge && make bridge

FROM debian:buster

RUN apt-get update  && apt-get install -y --no-install-recommends ca-certificates
RUN update-ca-certificates

WORKDIR "/opt"

ENV CONFIG_PATH ''
ENV VERBOSITY 3
ENV NUMBEROFWORKERS 1024

ENV LISTENERS__RONIN__CONFIG__RPCURL ''
ENV LISTENERS__RONIN__CONFIG__SECRET__BRIDGEOPERATOR__PLAINPRIVATEKEY ''
ENV LISTENERS__RONIN__CONFIG__TASKINTERVAL ''
ENV LISTENERS__RONIN__CONFIG__TRANSACTIONCHECKPERIOD ''
ENV LISTENERS__RONIN__CONFIG__FROMHEIGHT ''
ENV LISTENERS__RONIN__CONFIG__GASLIMITBUMPRATIO ''

ENV LISTENERS__RONIN__CONFIG__MAXTASKQUERY ''
ENV LISTENERS__RONIN__CONFIG__MAXPROCESSINGTASKS ''

ENV LISTENERS__RONIN__STATS__NODE ''
ENV LISTENERS__RONIN__STATS__HOST ''
ENV LISTENERS__RONIN__STATS__SECRET ''

ENV LISTENERS__ETHEREUM__CONFIG__RPCURL ''
ENV LISTENERS__ETHEREUM__CONFIG__SECRET__BRIDGEOPERATOR__PLAINPRIVATEKEY ''
ENV LISTENERS__ETHEREUM__CONFIG__GETLOGSBATCHSIZE ''

ENV DB__HOST ''
ENV DB__PORT ''
ENV DB__NAME ''
ENV DB__USERNAME ''
ENV DB__PASSWORD ''
ENV DB__CONNMAXLIFETIME ''
ENV DB__MAXIDLECONNS ''
ENV DB__MAXOPENCONNS ''

ENV VRF__WAITFORBLOCK ''
ENV VRF__CONTRACTADDRESS ''
ENV VRF__CONTRACTNAME ''
ENV VRF__SECRETKEY ''

COPY --from=builder /go/bin/bridge /usr/local/bin/bridge
COPY --from=builder /opt/bridge/config/ ./
COPY --from=builder /opt/bridge/docker/entrypoint.sh ./

ENTRYPOINT ["./entrypoint.sh"]
