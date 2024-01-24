FROM golang:1.21.6-alpine as builder

WORKDIR /app

COPY ./ ./

RUN go mod tidy
RUN go build -o ./bin/xyzApp ./main.go

FROM alpine:3

WORKDIR /app

RUN mkdir log
COPY --from=builder /app/config.json ./
COPY --from=builder /app/bin/xyzApp ./

EXPOSE 5005

CMD ./xyzApp
