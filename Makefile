.DEFAULT_GOAL := help

REPO := devopsgig

IMAGE_TEST_NAME := arithproducer_test
IMAGE_PROD_NAME := arithproducer

CONTAINER_TEST_NAME := "${IMAGE_TEST_NAME}"
CONTAINER_PROD_NAME := "${IMAGE_PROD_NAME}"

.PHONY: help build/test/image test build/prod/image run clen

help:
	@echo "------------------------------------------------------------------------"
	@echo "devopsgig arithproducer"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build/test/image: ## build test image
	@docker build -t "${REPO}"/"${IMAGE_TEST_NAME}" -f ./resources/test/Dockerfile .

build/prod/image: ## build prod image
	@docker build -t  "${REPO}"/"${IMAGE_PROD_NAME}" -f ./resources/prod/Dockerfile .

test: build/test/image ## run unit tests
	@docker run -it --rm --name "${CONTAINER_TEST_NAME}" "${REPO}"/"${IMAGE_TEST_NAME}"
	@printf "Removing "${REPO}"/"${IMAGE_TEST_NAME}" image\n\n"
	@docker rmi "${REPO}"/"${IMAGE_TEST_NAME}"

run: build/prod/image ## start the task generation and send each task as POST request to the API
	@docker run -d --name "${CONTAINER_PROD_NAME}" "${REPO}"/"${IMAGE_PROD_NAME}"

clean: ## stop and remove running production container
	@./scripts/rm-container.sh "${REPO}"/"${IMAGE_PROD_NAME}" "${CONTAINER_PROD_NAME}" &> /dev/null

