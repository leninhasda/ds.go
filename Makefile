default: run
	
run: build
	ds.go

build:
	go install
