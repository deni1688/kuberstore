mvn_build:
	./mvnw package
docker_build:
	docker build -t deni1688/warehouseserver . && rm -rf target
docker_push:
	docker push deni1688/warehouseserver:latest
