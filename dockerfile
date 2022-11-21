ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -o main .

FROM alpine:latest

RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
ENV GIN_MODE=release

EXPOSE 8080

CMD ["./main"]