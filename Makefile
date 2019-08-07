APP = clustery
BIN = $(APP)
BR = `git name-rev --name-only HEAD`
VER = `git describe --tags --abbrev=0`
COMMIT =`git rev-parse --short HEAD`
TIMESTM = `date -u '+%Y-%m-%d_%H:%M:%S%p'`
FORMAT = v$(VER)-$(COMMIT)-$(TIMESTM)
DOCTAG = $(VER)-$(BR)

.PHONY: info
info:
	make -v
	sudo docker version --format 'Client: {{ .Client.Version}} Server: {{ .Server.Version }}'
	godep version
	go version
	git describe --tags
	echo "namespace:"$(PRJ) "appname:"$(APP) "binary-name:"$(BIN) "version:"$(FORMAT)

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(BIN)