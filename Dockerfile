FROM golang AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -ldflags "-s -w -X 'microblog/common.Version=$(cat VERSION)' -extldflags '-static'" -o microblog

FROM alpine

RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates 2>/dev/null || true
ENV PORT=3000
COPY --from=builder /build/microblog /
EXPOSE 3000
WORKDIR /data
ENTRYPOINT ["/microblog"]
