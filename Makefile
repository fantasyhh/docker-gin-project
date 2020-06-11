# for docker
shell-nginx:
	docker exec -ti nginx_c /bin/sh

shell-web:
	docker exec -ti web_c /bin/sh

shell-db:
	docker exec -ti db_c /bin/sh



# for self
build:
	@go build -v .

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf drizzle
	go clean -i .
