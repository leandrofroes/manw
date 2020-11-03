GOCMD=go

.PHONY: build clean

all: build

build: 
	$(GOCMD) build manw

clean:
	rm -f manw
