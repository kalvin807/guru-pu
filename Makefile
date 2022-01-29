.phony: dev docker

dev:
	go run . -t "OTM0NzE4NTk2MDg3Njc2OTM4.Ye0Khw.PQEyJaIQcZoeZC7Hi-Z8VnBRRj0"

docker:
	docker build --tag gurupu .
