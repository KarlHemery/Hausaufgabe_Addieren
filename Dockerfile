FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go get -d -v ./...

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]