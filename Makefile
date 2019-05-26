.PHONY: all clean linux windows ship format archives
gob = go build
files = main.go $(wildcard ./**/*.go)
toplevel = main.go common Makefile README.md recv send setup

dist = dist
srvdist = /var/www/blog/down
uncom = $(dist)/uncompressed
com = $(dist)/compressed
source = $(dist)/source
tarball = $(source)/mastodial.tar
srczip = $(source)/mastodial.zip
compressed_prefix = cmp_

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

release: $(foreach X, $(allexes) $(tarball) $(tarball).gz $(srczip), $(srvdist)/$X) 

$(srvdist)/$(dist)/%: $(dist)/%
	mkdir -p $(dir $@) && cp $< $@

ship: $(compressedexes) archives

archives: $(tarball) $(tarball).gz $(srczip)

$(tarball): $(files) $(toplevel)
	mkdir -p $(source)
	tar cvf $(tarball) $(toplevel)

$(srczip): $(files) $(toplevel)
	mkdir -p $(source)
	zip -r $(srczip) $(toplevel)

$(com)/$(compressed_prefix)%: $(uncom)/%
	mkdir -p $(com)
	rm -f $@ #upx won't overwrite an existing file
	./upx --brute $< -o $@
	touch $@ #upx truncates the lower bits of the timestamp, fix that

%.gz: %
	gzip -k -f $<

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
