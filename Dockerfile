FROM golang:latest AS builder

WORKDIR /shortly

COPY go.mod .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o shortly cmd/*

FROM alpine:latest  

WORKDIR /root/
COPY --from=builder /shortly/shortly .

EXPOSE 8080
CMD ["./shortly"]