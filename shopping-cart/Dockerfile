# syntax=docker/dockerfile:1
FROM golang:1.16-alpine AS build
WORKDIR /go/cache
ADD go.mod ./
ADD go.sum ./
RUN go mod download

WORKDIR /go/release
ADD . ./
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app main.go

#scratch: 空的基礎映象，最小的基礎映象
#busybox: 帶一些常用的工具，方便除錯， 以及它的一些擴充套件busybox:glibc
#alpine: 另一個常用的基礎映象，帶包管理功能，方便下載其它依賴的包
FROM alpine as prod
#COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
#COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/release/app /
#COPY --from=build /go/release/config/config.json /etc/config/

EXPOSE 8809
CMD ["/app"]