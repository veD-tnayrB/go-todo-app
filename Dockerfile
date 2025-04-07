FROM golang:1.24.1

WORKDIR /app

COPY . /app

RUN go mod tidy
RUN go build -o ./bin/main ./cmd/app/main.go 

CMD ["./bin/main"] 
EXPOSE 5002 