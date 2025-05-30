FROM golang:1.24.1-alpine3.20 AS build

WORKDIR /app

COPY go.mod go.sum ./

ENV GODEBUG=netdns=go
RUN go mod download

COPY . .

RUN go build -o /app/gmail *.go

FROM alpine:3.19

WORKDIR /app

COPY database/migrations ./database/migrations/

COPY --from=build /app/gmail /app/gmail

EXPOSE 8080

CMD ["/app/gmail"]