gob := go build
files := root.go $(wildcard ./**/*.go)
dist := dist
native: $(files)
	$(gob) 

.PHONY: all clean linux windows strip


all: linux windows strip

clean: 
	go clean -i
	rm -r $(dist)


arch := 32 64
prefix := dist/mastodial-

distflags := -ldflags="-s -w"

strip:
	strip $(dist)/*linux*
	./upx --brute $(dist)/*

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
