FROM golang:1.20.3-alpine3.17 as builder

WORKDIR /app

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN go env -w GOPROXY=https://goproxy.cn && go env -w CGO_ENABLED=0

RUN apk add make gcc bash git

COPY . .

RUN go mod download

RUN make build

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/bin/fiber-starter .

EXPOSE 3000

CMD [ "./fiber-starter", "server"]