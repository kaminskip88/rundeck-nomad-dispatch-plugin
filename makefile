build: clean
	@mkdir -p ./build/nomad-dispatch-plugin/contents
	@GOOS=linux GOARCH=amd64 \
		go build -o ./build/nomad-dispatch-plugin/contents/nomad-dispatch-plugin
	@cp -r ./resources ./build/nomad-dispatch-plugin
	@cp ./plugin.yaml ./build/nomad-dispatch-plugin
	@cd ./build && zip -r nomad-dispatch-plugin.zip nomad-dispatch-plugin
clean: down
	@rm -rf ./build

up: build
	@docker-compose up -d

down:
	@docker-compose down

sh:
	@docker-compose exec rundeck /bin/bash

logs:
	@docker-compose logs rundeck