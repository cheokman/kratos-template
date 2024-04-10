FROM golang:1.22.1 as builder

ARG APP_RELATIVE_PATH
ENV GOOS="linux"
ENV GOARCH="386"
COPY . /src

WORKDIR /src/app/${APP_RELATIVE_PATH}

RUN make build

FROM debian:stable-slim

ARG APP_RELATIVE_PATH

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app/bin
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/configs /app/configs
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/migration /app/migration
COPY --from=builder /src/bin/migrate.linux-amd64 /app/bin/migrate

WORKDIR /app

CMD ["./bin/server", "-conf", "configs/"]