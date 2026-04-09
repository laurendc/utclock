.PHONY: build app run clean build-linux install-linux

BINARY  = utclock
APP     = UTClock.app
BUNDLE  = $(APP)/Contents/MacOS

build:
	go build -o $(BINARY) .

app: build
	mkdir -p $(BUNDLE)
	cp $(BINARY) $(BUNDLE)/$(BINARY)
	cp Info.plist $(APP)/Contents/Info.plist
	open $(APP)

run: app

build-linux:
	GOOS=linux go build -o $(BINARY) .

install-linux: build-linux
	install -m755 $(BINARY) ~/.local/bin/$(BINARY)

clean:
	rm -rf $(BINARY) $(APP)
