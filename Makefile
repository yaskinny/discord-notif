.PHONY: build

build:
	@CGO_ENABLED=0 go build -o bin/discord-notif cmd/main.go && \
		echo 'Built successfully'

install: build
	@cp bin/discord-notif /usr/local/bin/discord-notif && rm -rf ./bin

docker: build
	@docker build -t discord-notif:`git log --format='%h'` . && \
		echo "Docker image built"
