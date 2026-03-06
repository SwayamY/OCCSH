FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o occsh

EXPOSE 8080

CMD ["./occsh"]
