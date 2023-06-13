FROM golang:1.19-alpine as builder

COPY . /github.com/muratovdias/pocket-tg/
WORKDIR /github.com/muratovdias/pocket-tg/

RUN go mod download
RUN go build -o /bin/ cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/muratovdias/pocket-tg/bit/bot .
COPY --from=0 /github.com/muratovdias/pocket-tg/configs configs/

EXPOSE 8001

CMD["./bot"]
