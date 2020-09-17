build:
	go build -o bin/slsui cmd/slsui/*
	go build -o bin/downloadui cmd/downloadui/*
downloadui:
	go run ./cmd/downloadui