FROM golang:1.13-alpine as builder

WORKDIR /app

RUN apk --update add ca-certificates

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/server/*.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/app /app

ENTRYPOINT ["/app"]