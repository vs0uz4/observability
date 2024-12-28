FROM alpine:3.21

RUN apk --no-cache add ca-certificates bash

ARG USER_UID=10001
RUN adduser -D -u ${USER_UID} oteluser

USER ${USER_UID}

COPY --from=otel/opentelemetry-collector:latest /otelcol /otelcol

COPY config.yaml /etc/otelcol/config.yaml

ENTRYPOINT ["/otelcol"]
CMD ["--config", "/etc/otelcol/config.yaml"]

EXPOSE 4317 4318 55678 55679
