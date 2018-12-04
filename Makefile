HUB := zackbutcher
TAG := v0.1

default: build

build:
	go build .

echo-server.linux:
	GOOS=linux go build -a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo -o echo-server .

docker.build: echo-server.linux
	docker build -t ${HUB}/echo-server:${TAG} -f Dockerfile .

docker.run: docker.build
	docker run ${HUB}/echo-server:${TAG}

docker.push: docker.build
	docker push ${HUB}/echo-server:${TAG}