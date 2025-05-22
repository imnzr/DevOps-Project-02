FROM golang:1.24

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

EXPOSE 3000

CMD ["./app"]
