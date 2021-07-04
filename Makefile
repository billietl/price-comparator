test:
	FIRESTORE_EMULATOR_HOST=127.0.0.1:8081 go test ./...

firestore:
	gcloud beta emulators firestore start --host-port=127.0.0.1:8081
