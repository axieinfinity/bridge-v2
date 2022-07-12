# Build Geth in a stock Go builder container
FROM golang:1.17.0-alpine3.13 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

COPY . /opt
RUN cd /opt && make bridge

FROM alpine:3.15

RUN apk add --no-cache ca-certificates
WORKDIR "/opt"

ENV CONFIG_PATH ''
ENV RONIN_RPC ''
ENV RONIN_VALIDATOR_KEY ''
ENV RONIN_RELAYER_KEY ''
ENV RONIN_TASK_INTERVAL ''
ENV RONIN_TRANSACTION_CHECK_PERIOD ''
ENV RONIN_MAX_PROCESSING_TASKS ''

ENV RONIN_MAX_TASK_QUERY ''

ENV ETHEREUM_RPC ''
ENV ETHEREUM_VALIDATOR_KEY ''
ENV ETHEREUM_RELAYER_KEY ''
ENV ETHEREUM_GET_LOGS_BATCH_SIZE ''
ENV VERBOSITY 3
ENV NUMBER_OF_WORKERS 1024

ENV DB_HOST ''
ENV DB_PORT ''
ENV DB_NAME ''
ENV DB_USERNAME ''
ENV DB_PASSWORD ''
ENV DB_CONN_MAX_LIFE_TIME ''
ENV DB_MAX_IDLE_CONNS ''
ENV DB_MAX_OPEN_CONNS ''

ENV CLEANER ''

COPY --from=builder /go/bin/bridge /usr/local/bin/bridge
COPY --from=builder /opt/config/ ./
COPY --from=builder /opt/contracts ./contracts
COPY --from=builder /opt/docker/entrypoint.sh ./

ENTRYPOINT ["./entrypoint.sh"]