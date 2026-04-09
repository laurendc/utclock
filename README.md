# utclock

A UTC clock for your system bar — macOS menubar app or Linux i3bar block.

## Requirements

- Go 1.21+
- macOS or Linux with i3/i3blocks

## Build & Run

```bash
make app
```

This builds the binary, creates `UTClock.app`, and launches it. The current UTC time (`HH:MM UTC`) appears in your menu bar. Click it to see a dropdown with both your local time and UTC updated every second.

## Auto-start on login

1. Open **System Settings → General → Login Items**
2. Click **+** and add `UTClock.app` from this directory

## Linux (i3blocks)

Build and install the binary to `~/.local/bin/`:

```bash
make install-linux
```

Then add this block to `~/.config/i3blocks/config` (before the `[time]` block so UTC appears to its left):

```ini
# ----- UTC Clock -----
[utclock]
command=~/.local/bin/utclock
interval=1
markup=pango
```

Reload i3 with `$mod+Shift+r` — the UTC time appears in your i3bar styled in Nord blue (`#88C0D0`).

## Development

```bash
make build         # compile binary only (macOS)
make build-linux   # compile binary only (Linux)
make clean         # remove binary and .app bundle
```
