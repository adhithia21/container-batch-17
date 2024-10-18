FROM golang:1.20-alpine3.19 as builder
WORKDIR /app
COPY . .
RUN go build -o myapp

FROM alpine:latest as release
WORKDIR /app
COPY --from=builder /app/myapp .
EXPOSE 8000
ENTRYPOINT ["./myapp"]


