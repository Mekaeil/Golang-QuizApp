FROM golang:1.20-alpine as builder

ENV GOBINARIES /go/bin
ENV PATH="/usr/local/go/bin:${PATH}"

RUN apk add --no-cache git
RUN go install github.com/cespare/reflex@latest


FROM alpine

RUN apk add --update --no-cache go vim git make musl-dev curl

COPY --from=builder /go/bin/reflex /usr/bin/reflex

RUN apk add --no-cache ca-certificates