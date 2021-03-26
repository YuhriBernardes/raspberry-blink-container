FROM golang:1.16.2-buster as build

WORKDIR /build
COPY ./ ./

RUN go mod tidy

ENV GOOS=linux
ENV GOARCH=arm
ENV GOARM=7

RUN go build -o=blinker ./main.go

FROM arm32v7/alpine:3.10

COPY --from=build /build/blinker ./

# RUN chmod +x ./blinker

ENTRYPOINT ["./blinker"]