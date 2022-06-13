FROM golang:1.18-alpine as builder

RUN apk add --no-cache git make
ENV GOOS=linux
ENV CGO_ENABLED=0
WORKDIR /src
COPY go.sum go.sum
COPY go.mod go.mod
RUN go mod download
COPY . .
RUN go build -a -installsuffix cgo -o nada-datasetenricher ./cmd/nada-datasetenricher

FROM alpine:3
WORKDIR /app
COPY --from=builder /src/nada-datasetenricher /app/nada-datasetenricher
ENV NADA_BACKEND_URL http://localhost:8080/api/query
CMD ["/app/nada-datasetenricher"]
