.DEFAULT_GOAL := help

## -- Section Delimiter --

## This help message
## Which can also be multiline
.PHONY: help
help:
	@printf "Usage\n";

	@awk '{ \
			if ($$0 ~ /^.PHONY: [a-zA-Z\-\_0-9]+$$/) { \
				helpCommand = substr($$0, index($$0, ":") + 2); \
				if (helpMessage) { \
					printf "\033[36m%-20s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^[a-zA-Z\-\_0-9.]+:/) { \
				helpCommand = substr($$0, 0, index($$0, ":")); \
				if (helpMessage) { \
					printf "\033[36m%-20s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^##/) { \
				if (helpMessage) { \
					helpMessage = helpMessage"\n                     "substr($$0, 3); \
				} else { \
					helpMessage = substr($$0, 3); \
				} \
			} else { \
				if (helpMessage) { \
					print "\n                     "helpMessage"\n" \
				} \
				helpMessage = ""; \
			} \
		}' \
		$(MAKEFILE_LIST)


###################################################################################################
#########################                  INSTALLING                     #########################
###################################################################################################
APP_DOCKER_NAME = "quiz-app"
INSTALLING_MESSAGE = "🚀🚀 Application is ready to GO! 🥳🥳 \n"
CLEAR_CACHE_MESSAGE = "🧹🧹 Let's GO to clear all of caches! 🥳🥳 Now the  🗑️  is empty! 😅😅 \n"
TEST_MESSAGE = "🧪🧪 Let's GO to test the application! \n"
WATCH_MESSAGE = "👀👀 Let's watching your changes... \n"

## INSTALLING AND RUNNING DOCKER
.PHONY: install
install:
	cp .env.example .env
	cd deployment && docker-compose up -d --force-recreate && docker-compose build --force-rm
	@echo $(INSTALLING_MESSAGE)
	make watch
	exit 0


## BUILDING THE APPLICATION
.PHONY: rebuild
rebuild:
	make install

## WATCHING
.PHONY: watch
watch:
	@echo $(WATCH_MESSAGE)
	docker exec -it $(APP_DOCKER_NAME) reflex -r '(\.go)' -s -- sh -c 'go run bootstrap/main.go'
	exit 0


###################################################################################################
#########################                USEFUL COMMANDS                  #########################
###################################################################################################

## Clear Cache
.PHONY: clean
clean:
	docker exec -it $(APP_DOCKER_NAME) go clean
	docker exec -it $(APP_DOCKER_NAME) go clean -cache
	docker exec -it $(APP_DOCKER_NAME) go clean -testcache
	@echo $(CLEAR_CACHE_MESSAGE)
	exit 0

## RUNNING TESTCASES
.PHONY: test
test:
	@echo $(TEST_MESSAGE)
	go test ./test/... -v -failfast -shuffle=on -benchmem -coverprofile cover.out
	@echo "\n"
	exit 0