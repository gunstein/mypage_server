FROM golang:latest as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o mypage_server -ldflags="-w -s"

#FROM scratch
FROM centos:latest
COPY --from=builder /app/mypage_server /

# Copy CA certificates to prevent x509: certificate signed by unknown authority errors
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 3000
ENTRYPOINT ["/mypage_server"]