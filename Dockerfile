FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache nodejs npm

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cd web && npm install && npm run build

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /am ./cmd/server

FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /am /am
COPY configs/config.yaml /configs/config.yaml

EXPOSE 8080

ENTRYPOINT ["/am"]
