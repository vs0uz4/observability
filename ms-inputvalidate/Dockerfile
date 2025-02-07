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
    addgroup --system inputvalidate && \ 
    adduser --system --ingroup inputvalidate inputvalidate

WORKDIR /app

COPY ./pkg /app/pkg
COPY ./ms-inputvalidate /app/ms-inputvalidate

WORKDIR /app/ms-inputvalidate

RUN go mod download && \
    mv /app/ms-inputvalidate/cmd/inputvalidate/.env /app/ms-inputvalidate/.env && \
	CGO_ENABLED=0 GOOS=linux go build -o inputvalidateapp ./cmd/inputvalidate

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
USER inputvalidate

COPY --from=builder /app/ms-inputvalidate/.env .
COPY --from=builder /app/ms-inputvalidate/inputvalidateapp .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/America/Sao_Paulo /usr/share/zoneinfo/America/Sao_Paulo

ENV TZ=America/Sao_Paulo

ENTRYPOINT [ "./inputvalidateapp" ]