DOCKER_IMAGE_NAME=slide-decks

.PHONY: build
build:
	docker image build -f Dockerfile -t ${DOCKER_IMAGE_NAME} .

.PHONY: run
run:
	docker run --rm -it --name slide-decks \
		-v `pwd`:/go/src/github.com/task4233/slide-decks \
		-p 3000:3000 \
		${DOCKER_IMAGE_NAME}

