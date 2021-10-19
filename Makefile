build_all:
	cd admin; make go_build && make docker_build && make docker_push & cd ..
	cd catalog; make go_build && make docker_build && make docker_push & cd ..
	cd warehouse; make mvn_build && make docker_build && make docker_push & cd ..
	cd ui; make ts_build && make docker_build && make docker_push & cd ..
