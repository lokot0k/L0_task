FROM golang:1.24
WORKDIR /L0_task
COPY ../.. /L0_task
RUN go mod download
RUN go build -o L0_task_app ./cmd/web_app
EXPOSE 8080
CMD ["./L0_task_app"]
