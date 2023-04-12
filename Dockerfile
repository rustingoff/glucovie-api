FROM golang:1.19

WORKDIR /app/go-app

COPY . .

RUN go mod download
RUN go mod tidy

RUN go build cmd/main.go

EXPOSE 8000

CMD ["./main"]