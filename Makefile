build:
	go build -i -v -o ./cmd/serve . ;\
	echo "\033[0;32m[successed]\033[0m" $<

linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./cmd/serve . ;\
	echo "\033[0;32m[successed]\033[0m" $<

serve:
	cd ./cmd ;\
	./serve

run: serve

clean:
	sudo rm -rf ./cmd/*
	echo "ok"

test:
	sudo go test -v -run ./ ./