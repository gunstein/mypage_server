FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN CGO_ENABLED=0 go build -o mypage_server -ldflags="-w -s"

#FROM scratch
FROM centos:latest
COPY --from=builder /app/mypage_server /
COPY --from=builder /app/gunstein_vatnar_no /gunstein_vatnar_no

# Copy CA certificates to prevent x509: certificate signed by unknown authority errors
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 3000
ENTRYPOINT ["/mypage_server"]