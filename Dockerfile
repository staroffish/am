FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache nodejs npm

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cd web && npm install && npm run build

# SQLite 驱动是纯 Go 实现，不需要 CGO
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /am ./cmd/server

FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /am /am
COPY configs/config.yaml /app/configs/config.yaml

# SQLite 数据文件目录：必须挂到本机磁盘卷上，不能是网络共享盘
VOLUME ["/app/data"]

EXPOSE 8222

ENTRYPOINT ["/am", "-config", "/app/configs/config.yaml"]
