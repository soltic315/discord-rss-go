FROM golang:latest as builder 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -o main .

FROM alpine:latest

COPY --from=builder  /app/main .

EXPOSE 8080

ENTRYPOINT ["./main"]