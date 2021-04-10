DOCKER_IMAGE_NAME=slide-decks
PORT=3000

.PHONY: build
build:
	docker image build -f Dockerfile -t ${DOCKER_IMAGE_NAME} .

.PHONY: run
run:
	docker run --rm -it --name slide-decks \
		-v `pwd`:/go/src/github.com/task4233/slide-decks \
		-p ${PORT}:${PORT} \
		${DOCKER_IMAGE_NAME}

.PHONY: run-local
run-local:
	present -http "0.0.0.0:${PORT}" -orighost localhost