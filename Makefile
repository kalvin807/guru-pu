.phony: dev docker

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

dev:
	go run . -token "${DISCORD_BOT_TOKEN}" -redis "${REDIS_URL}"

docker:
	docker build --tag gurupu --build-arg DISCORD_BOT_TOKEN="${DISCORD_BOT_TOKEN}" --build-arg REDIS_URL="${REDIS_URL}" .
