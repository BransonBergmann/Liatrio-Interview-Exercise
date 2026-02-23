# syntax=docker/dockerfile:experimental

FROM golang:1.25.7

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o send_message .

EXPOSE 80

CMD [ "./send_message" ]