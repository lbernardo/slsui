build:
	env GOOS=darwin go build -o bin/darwin/slsui cmd/slsui/*
	env GOOS=linux go build -o bin/linux/slsui cmd/slsui/*
downloadui:
	go run ./cmd/downloadui