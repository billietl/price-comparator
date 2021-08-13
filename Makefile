run:
	FIRESTORE_EMULATOR_HOST=127.0.0.1:8081 go run price-comparator.go

build:
	go build -o target/price-comparator price-comparator.go

test:
	FIRESTORE_EMULATOR_HOST=127.0.0.1:8081 go test -v ./...

firestore:
	gcloud beta emulators firestore start --host-port=127.0.0.1:8081
