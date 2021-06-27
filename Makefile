test:
	gcloud beta emulators firestore start --host-port=127.0.0.1:8888 &
	pid=$$!
	FIRESTORE_EMULATOR_HOST=127.0.0.1:8888 go test ./...
	kill $$pid
