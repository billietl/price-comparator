run:
	GOOGLE_PROJECT_ID=price-comparator-dev go run price-comparator.go

build:
	go build -o target/price-comparator price-comparator.go

test: format
	FIRESTORE_EMULATOR_HOST=127.0.0.1:8081 go test -v ./...

firestore:
	gcloud beta emulators firestore start --host-port=127.0.0.1:8081

lint:
	golint ./dao ./model ./web
	golint price-comparator.go

superlint:
	docker run -it --rm -v $$PWD:/tmp/lint -e RUN_LOCAL=true github/super-linter:v4 | tee target/superlint.log

format:
	goimports -w ./dao ./model ./web ./price-comparator.go
	go vet ./...
	go fmt ./...
