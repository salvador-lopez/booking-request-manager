SHELL := /bin/bash

tests:
	@echo Running unit and e2e tests
	docker-compose up go_tests

server-run:
	@echo Running the api server
	docker-compose up booking_request_manager