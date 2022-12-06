FROM golang AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -ldflags "-s -w -extldflags '-static'" -o microblog

FROM alpine

ENV PORT=3000
COPY --from=builder /build/microblog /
EXPOSE 3000
WORKDIR /data
ENTRYPOINT ["/microblog"]
