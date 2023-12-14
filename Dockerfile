FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]