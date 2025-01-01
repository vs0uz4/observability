FROM golang:1.23.3 AS builder

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get dist-upgrade -y && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    echo "America/Sao_Paulo" > /etc/timezone && \
    ln -sf /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    addgroup --system weatherzip && \ 
    adduser --system --ingroup weatherzip weatherzip

WORKDIR /app

COPY ./pkg /app/pkg
COPY ./ms-weatherzip /app/ms-weatherzip

WORKDIR /app/ms-weatherzip

RUN go mod download && \
    mv /app/ms-weatherzip/cmd/weatherzip/.env /app/ms-weatherzip/.env && \
	CGO_ENABLED=0 GOOS=linux go build -o weatherzipapp ./cmd/weatherzip

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
USER weatherzip

COPY --from=builder /app/ms-weatherzip/.env .
COPY --from=builder /app/ms-weatherzip/weatherzipapp .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/America/Sao_Paulo /usr/share/zoneinfo/America/Sao_Paulo

ENV TZ=America/Sao_Paulo

ENTRYPOINT [ "./weatherzipapp" ]