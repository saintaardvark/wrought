run: wrought
	./wrought

wrought: *.go */*.go
	go build .
