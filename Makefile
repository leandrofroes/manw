GOCMD=go

.PHONY: build clean

all: build

build: 
	$(GOCMD) build

clean:
	rm -f manw
