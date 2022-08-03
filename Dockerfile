# syntax=docker/dockerfile:1

FROM golang:alpine as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o mindustry-bp

FROM alpine:3.16.1
WORKDIR /app
COPY --from=builder ./app/mindustry-bp ./mindustry-bp
COPY  ./templates ./templates
COPY  ./static ./static
COPY config.yml ./config.yml
EXPOSE 8080
USER nobody:nobody
CMD [ "./mindustry-bp" ]
