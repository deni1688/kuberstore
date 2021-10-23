go_build:
	go build -tags netgo -o app
docker_build:
	docker build -t deni1688/publisherserver . && rm app
docker_push:
	docker push deni1688/publisherserver:latest
