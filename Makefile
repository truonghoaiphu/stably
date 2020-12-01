build:
	go build -i -v -o ./cmd/serve . ;\
	echo "\033[0;32m[successed]\033[0m" $<

linux:
	rm -rf go.mod
	rm -rf go.sum
	echo "module stably" >> go.mod
	echo "go 1.14" >> go.mod
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./cmd/serve . ;\
	echo "\033[0;32m[successed]\033[0m" $<

serve:
	cd ./cmd ;\
	./serve

run: serve

up:
	cd ./deployment ;\
	docker-compose up -d
	
down:
	cd ./deployment ;\
	docker-compose down

test:
	go test -v ./... -cover

clean:
	sudo rm -rf ./cmd/*
	echo "ok"

main_test:
	sudo go test -v -run ./ ./