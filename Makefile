BINNAME=llb
BUILDDIR=release

all:
	go get github.com/mitchellh/gox
	go get
	gox --output=$(BUILDDIR)"/{{.OS}}/{{.Arch}}/"$(BINNAME)
	build/prepare.sh

clean: 
	rm -rf ./$(BUILDDIR)/
