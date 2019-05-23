.PHONY: all clean linux windows ship format tarball
gob = go build
files = main.go $(wildcard ./**/*.go)
toplevel = main.go common Makefile README.md recv send setup
dist = dist
uncom = $(dist)/uncompressed
com = $(dist)/compressed
tarball = $(dist)/mastodial.tar
compressed_prefix = cmp_

exes = mastodial-linux32 mastodial-linux64 mastodial-windows32.exe mastodial-windows64.exe

native: $(files)
	$(gob) 

format:
	gofmt -l -s -w .


all: linux windows

clean: 
	go clean -i
	rm -r $(dist)

tarball: $(tarball)

ship: $(foreach X, $(exes), $(com)/$(compressed_prefix)$X) $(tarball) $(tarball).gz

$(dist):
	mkdir -p $(dist)

$(tarball): $(files) $(dist)
	tar cvf $(tarball) $(toplevel)

$(com)/$(compressed_prefix)%: $(uncom)/%
	mkdir -p $(com)
	cp $< $@
	strip $@
	./upx --brute $@

%.gz: %
	gzip -k -f $<

arch := 32 64
prefix := $(uncom)/mastodial-

distflags := -ldflags="-s -w"
linux: $(foreach X,$(arch),$(prefix)linux$X)

windows: $(foreach X,$(arch),$(prefix)windows$X.exe)

$(prefix)linux64: $(files)
	GOOS=linux GOARCH=amd64 $(gob) $(distflags) -o $@ 

$(prefix)linux32: $(files)
	GOOS=linux GOARCH=386 $(gob) $(distflags) -o $@ 

$(prefix)windows64.exe: $(files)
	GOOS=windows GOARCH=amd64 $(gob) $(distflags) -o $@ 

$(prefix)windows32.exe: $(files)
	GOOS=windows GOARCH=386 $(gob) $(distflags) -o $@ 
