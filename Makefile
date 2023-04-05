.PHONY: dockerize
dockerize:
	docker compose up -d --build

.PHONY: stop_db
stop_db:
	docker container stop user-service-mariadb-1 \
	&& docker wait user-service-mariadb-1 \
	&& docker container rm user-service-mariadb-1 \
	&& docker image rm mariadb

.PHONY: stop_service
stop_service:
	docker kill --signal SIGINT user-service-fiber-application-1 \
	&& docker wait user-service-fiber-application-1 \
	&& docker container rm user-service-fiber-application-1 \
	&& docker image rm user-service

.PHONE: generate_docs
generate_docs:
	~/go/bin/swag fmt \
	&& ~/go/bin/swag init -g ./internal/infrastructure/fiber.go

.PHONY: test_e2e
test_e2e:
	sh ./scripts/test-e2e.sh
	