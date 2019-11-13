FROM golang:latest as builder

WORKDIR /src

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /src/main .

CMD ["./main"]