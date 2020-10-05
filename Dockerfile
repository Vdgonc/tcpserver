FROM golang:1.14.4 AS builder
WORKDIR /go/src/github.com/vdgonc/tcpserver/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build  -o app .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/vdgonc/tcpserver .
CMD ["./app"]
