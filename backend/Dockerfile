FROM golang:1.18 AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0 
RUN go build -o main cmd/main.go

FROM alpine:latest
COPY --from=builder /app ./

EXPOSE 8080
ENTRYPOINT ["./main"]