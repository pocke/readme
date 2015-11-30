all:
	go get github.com/jteeuwen/go-bindata/...
	go-bindata readme.template
	go build
