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
    addgroup --system searchzip && \ 
    adduser --system --ingroup searchzip searchzip

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd      /app/cmd
COPY ./config   /app/config
COPY ./internal /app/internal

RUN mv /app/cmd/searchzip/.env /app/.env && \
	CGO_ENABLED=0 GOOS=linux go build -o searchzipapp ./cmd/searchzip

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
USER searchzip

COPY --from=builder /app/.env .
COPY --from=builder /app/searchzipapp .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/America/Sao_Paulo /usr/share/zoneinfo/America/Sao_Paulo

ENV TZ=America/Sao_Paulo

ENTRYPOINT [ "./searchzipapp" ]