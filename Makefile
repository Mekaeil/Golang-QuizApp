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


INSTALLING_MESSAGE = "🚀🚀 Application is ready to GO!"
WATCH_MESSAGE = "👀👀 Let's watching your changes..."

## INSTALLING AND RUNNING DOCKER
.PHONY: install
install:
	cp .env.example .env
	cd deployment && docker-compose up -d --force-recreate && docker-compose build --force-rm
	@echo $(INSTALLING_MESSAGE)
	exit 0


## RUN THE APPLICATION
.PHONY: run
run: 
	go run bootstrap/main.go 
	exit 0
