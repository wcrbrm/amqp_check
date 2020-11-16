FROM golang:1.13-alpine
WORKDIR /app
ADD go.mod go.sum *.go /app/
RUN go build -o /app/amqp_check .

FROM alpine
WORKDIR /app
COPY --from=0 /app/amqp_check /app/amqp_check
ENTRYPOINT /app/amqp_check
