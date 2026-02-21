# syntax=docker/dockerfile:experimental

FROM golang:1.25.7

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build send_message

EXPOSE 8080

CMD [ "/build/send_message" ]