FROM golang:1.16.2-alpine3.13 as build

WORKDIR /build
COPY ./ ./

RUN go mod tidy

RUN go build -o=blinker ./main.go

FROM alpine:3

WORKDIR /app

COPY --from=build /build/blinker ./

RUN chmod +x ./blinker

ENTRYPOINT ["./blinker"]