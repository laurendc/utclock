# utclock

A macOS menubar app that shows the current UTC time alongside your local time — always visible in the top-right corner of your screen.

## Requirements

- Go 1.21+
- macOS

## Build & Run

```bash
make app
```

This builds the binary, creates `UTClock.app`, and launches it. The current UTC time (`HH:MM UTC`) appears in your menu bar. Click it to see a dropdown with both your local time and UTC updated every second.

## Auto-start on login

1. Open **System Settings → General → Login Items**
2. Click **+** and add `UTClock.app` from this directory

## Development

```bash
make build   # compile binary only
make clean   # remove binary and .app bundle
```
