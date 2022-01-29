FROM golang:1.17-alpine

WORKDIR app

COPY . .

RUN "go build -o ./bin/main ."

ARG DISCORD_BOT_TOKEN
ARG REDIS_URL

ENV DISCORD_BOT_TOKEN ${DISCORD_BOT_TOKEN}
ENV REDIS_URL ${REDIS_URL}

CMD ["./bin/main", "-token", ${DISCORD_BOT_TOKEN}, "-redis", ${REDIS_URL}]
