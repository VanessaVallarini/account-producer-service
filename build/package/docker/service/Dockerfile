FROM golang:1.18 as builder

WORKDIR /app
COPY cmd/account-producer-service .
COPY api /app/api
COPY internal /app/internal
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o account-producer-service

COPY --from=builder /app/account-producer-service /app/app/account-producer-service
EXPOSE 8080
ENTRYPOINT [ "/executor", "/app/app/account-producer-service" ]
