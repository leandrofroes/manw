GOCMD=go

.PHONY: build clean

all: build

build: manw.go 
	$(GOCMD) build $<

clean:
	rm -f manw
