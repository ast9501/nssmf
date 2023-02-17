FROM golang:1.19.5-alpine3.17 AS builder

WORKDIR /go/app
ADD . /go/app
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN apk add make
RUN make all

FROM alpine:3.17
WORKDIR /app
RUN mkdir -p config/TLS

# copy tls 
COPY --from=builder /go/app/config/TLS config/TLS

# copy config
COPY --from=builder /go/app/config/nssmf.env config/

# copy exec
COPY --from=builder /go/app/bin /app/

CMD ["./nssmf", "-c", "config"]