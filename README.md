# Mastodial

## Usage

Run the software without any arguments for now.

## Building

If you simply want to download and build the software for your own use, you only need the .go files and to do a `go build`. A simple `make` will do the same thing.

`make all` will build the four release packages (32 and 64 bit versions for Windows and Linux) and run strip(1) and [UPX](https://github.com/upx/upx) on them, if available. You probably won't have to do this unless you're distributing the binaries yourself, in which case, go for it.
