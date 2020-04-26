FROM golang:1.14 AS builder
WORKDIR /usr/src/inicom
COPY . /usr/src/inicom
RUN go mod download && \
    go build -v -o bin/inicom cmd/inicom/main.go && \
    ls -latr bin


FROM scratch
COPY --from=builder /usr/src/inicom/bin/inicom .
CMD ["./inicom"]
