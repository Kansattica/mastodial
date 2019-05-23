.PHONY: all clean linux windows ship format archives
gob = go build
files = main.go $(wildcard ./**/*.go)
toplevel = main.go common Makefile README.md recv send setup
dist = dist
uncom = $(dist)/uncompressed
com = $(dist)/compressed
source = $(dist)/source
tarball = $(source)/mastodial.tar
srczip = $(source)/mastodial.zip
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

ship: $(foreach X, $(exes), $(com)/$(compressed_prefix)$X) archives

archives: $(tarball) $(tarball).gz $(srczip)

$(tarball): $(files) $(toplevel)
	mkdir -p $(source)
	tar cvf $(tarball) $(toplevel)

$(srczip): $(files) $(toplevel)
	mkdir -p $(source)
	zip -r $(srczip) $(toplevel)

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
