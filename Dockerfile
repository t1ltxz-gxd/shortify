FROM golang:bullseye AS builder

WORKDIR /usr/local/src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/app cmd/grpc_server/main.go

FROM debian:bullseye-slim as runtime

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates && \
    apt-get install -y sudo && \
    apt-get autoremove -y && \
    apt-get clean

ARG UID=10001
RUN adduser \
--disabled-password \
--gecos "" \
--home "/nonexistent" \
--shell "/sbin/nologin" \
--no-create-home \
--uid "${UID}" \
shortify

RUN mkdir -p logs && chown shortify:shortify logs

COPY --from=builder /usr/local/src/bin/app /usr/bin

COPY config/config.yml /config/config.yml

COPY .env .

COPY docker-entrypoint.sh /usr/local/sbin/docker-entrypoint.sh
ENTRYPOINT ["/usr/local/sbin/docker-entrypoint.sh"]
CMD /usr/bin/app "$LOG_DIR"