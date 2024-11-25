FROM golang:1.22.3-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o stester ./cmd/stester

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/stester .

ENTRYPOINT ["/app/stester"]
