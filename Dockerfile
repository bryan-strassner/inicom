FROM golang:1.14 AS builder
WORKDIR /usr/src/inicom
COPY . /usr/src/inicom
RUN go mod download && \
    go build -o bin/inicom

FROM scratch
COPY --from=builder /go/src/inicom/bin/inicom .
CMD ["./inicom"]
