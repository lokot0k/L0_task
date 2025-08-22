FROM golang:1.24
WORKDIR /L0_task
COPY ../.. /L0_task
RUN go mod download
RUN go build -o L0_kafka_loader ./cmd/kafka_loader
CMD ["./L0_kafka_loader"]
