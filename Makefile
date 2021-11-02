SHELL := /bin/bash

init:
	@echo Building the containers needed to run, testing and benchmarking the application
	docker-compose build

tests:
	@echo Running unit and e2e tests
	docker-compose up go_tests

benchmark:
	@echo Running application benchmark
	docker-compose up go_benchmark

server-run:
	@echo Running the api server
	docker-compose up booking_request_manager