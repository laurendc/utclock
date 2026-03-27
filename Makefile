.PHONY: build app run clean

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

clean:
	rm -rf $(BINARY) $(APP)
