#Build
FROM golang:1.24.1-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run
FROM alpine:3.21 
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main" ]
