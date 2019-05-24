# Mastodial

## Usage

Running Mastodial without any arguments will display help information. Try `./mastodial`. If you downloaded one of the release binaries, feel free to rename it to something easier to type, like `mastodial` or `md`.

Be careful with your configuration file! Anyone with the authentication information inside can post to your Mastodon account.
## Building

If you simply want to download and build the software for your own use, you only need the .go files and to do a `go build`. A simple `make` will do the same thing. I have provided source tarballs and .zip files for those who would rather not clone a git repository over a slow connection.

`make all` will build four executables, 32 and 64 bit versions for Windows and Linux, respectively.

`make ship` will:
	- build the three source archives (tarball, gzipped tarball, and .zip)
	- build the four release packages (32 and 64 bit versions for Windows and Linux) and run [UPX](https://github.com/upx/upx) on them, if available. You probably won't have to do this unless you're distributing the binaries yourself.
