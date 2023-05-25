FROM golang:alpine as builder

ENV CGO_ENABLED 0
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
# RUN cd ./cmd/cron && go build -ldflags="-s -w" -o /app/cron.job
RUN cd ./cmd/api && go build -o /app/server.app


FROM alpine:latest

ENV TZ Asia/Shanghai

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk update --no-cache \
    && apk add --no-cache ca-certificates tzdata \
    && cp /usr/share/zoneinfo/${TZ} /etc/localtime \
    && rm -rf /tmp/*

WORKDIR /app
COPY --from=builder /app/server.app /app/server.app

CMD ["./server.app"]
