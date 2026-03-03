# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG IMAGE_TAG=unknown
ARG GIT_COMMIT=unknown

RUN CGO_ENABLED=0 go build \
  -trimpath \
  -ldflags="-s -w -X main.ImageTag=${IMAGE_TAG} -X main.GitCommit=${GIT_COMMIT}" \
  -o send_message .

FROM gcr.io/distroless/static-debian12

WORKDIR /app


COPY --from=builder /build/send_message .


EXPOSE 80

CMD [ "./send_message" ]