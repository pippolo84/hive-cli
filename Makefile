
.PHONY: build clean

build:
	go build -o cli

clean:
	rm -f ./cli