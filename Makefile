TAG=`git describe --tags`
REFHASH=`git log --pretty=format:'%h' -n1`
VERSION=$(TAG)-$(REFHASH)
run: wrought
	./wrought

wrought: *.go */*.go
	go build -ldflags "-X main.version=$(VERSION)" .

clean:
	rm wrought
