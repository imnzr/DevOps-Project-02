FROM golang:1.24 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# STAGE 2 

FROM alpine:3.14

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /root/

COPY --from=builder /usr/src/app/app .  

EXPOSE 3000

CMD ["./app"]
