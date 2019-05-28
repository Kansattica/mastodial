.PHONY: all clean linux windows format archives release ship
gob = go build
files = main.go go.mod go.sum $(wildcard ./**/*.go)
toplevel = main.go common Makefile README.md recv send setup go.mod go.sum

dist = dist
srvdist = /var/www/blog/down
uncom = $(dist)/uncompressed
com = $(dist)/compressed
source = $(dist)/source
tarball = $(source)/mastodial.tar
srczip = $(source)/mastodial.zip
compressed_prefix = cmp_

gzip = ./zopfli #gzip

exes = mastodial-linux32 mastodial-linux64 mastodial-windows32.exe mastodial-windows64.exe

native: format mastodial

mastodial: $(files)
	$(gob) 

format:
	gofmt -l -s -w .

all: linux windows

clean: 
	go clean -i
	rm -r $(dist)


uncompressedexes = $(foreach X, $(exes), $(uncom)/$X)
compressedexes = $(foreach X, $(exes), $(com)/$(compressed_prefix)$X)

allexes = $(uncompressedexes) $(compressedexes)

release: $(foreach X, $(allexes) $(tarball) $(srczip), $(srvdist)/$X $(srvdist)/$X.gz) 

$(srvdist)/$(dist)/%: $(dist)/%
	mkdir -p $(dir $@) && cp $< $@

ship: $(compressedexes) archives

archives: $(tarball) $(tarball).gz $(srczip)

$(tarball): $(files) $(toplevel)
	mkdir -p $(source)
	tar --exclude='.[^/]*'  -cvf $(tarball) $(toplevel)

$(srczip): $(files) $(toplevel)
	mkdir -p $(source)
	zip -r $(srczip) $(toplevel)

$(com)/$(compressed_prefix)%: $(uncom)/% upx
	mkdir -p $(com)
	rm -f $@ #upx won't overwrite an existing file
	./upx --brute $< -o $@
	touch $@ #upx truncates the lower bits of the timestamp, fix that

upx:
	wget https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz -O - | xzcat | tar xv
	mv upx-3.95-amd64_linux/upx .
	rm -r upx-3.95-amd64_linux

%.gz: % $(gzip)
	$(gzip) -k -f $<

gzip:

zopfli:
	git clone https://github.com/google/zopfli.git
	cd zopfli && make
	mv zopfli/zopfli ./z
	rm -rf zopfli
	mv z zopfli

arch := 32 64
prefix := $(uncom)/mastodial-

distflags := -ldflags="-s -w"
linux: $(foreach X,$(arch),$(prefix)linux$X)

windows: $(foreach X,$(arch),$(prefix)windows$X.exe)

%.exe: %
	mv $< $@ 

$(prefix)%64: $(files)
	GOOS=$* GOARCH=amd64 $(gob) $(distflags) -o $@ 

$(prefix)%32: $(files)
	GOOS=$* GOARCH=386 $(gob) $(distflags) -o $@ 
