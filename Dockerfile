# syntax=docker/dockerfile:1

# Build
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /gurupu


# Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build /gurupu /gurupu

ARG DISCORD_BOT_TOKEN
ARG REDIS_URL

ENV DISCORD_BOT_TOKEN=$DISCORD_BOT_TOKEN
ENV REDIS_URL=$REDIS_URL

CMD /gurupu -token ${DISCORD_BOT_TOKEN} -redis ${REDIS_URL}
