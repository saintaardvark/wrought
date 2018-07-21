run: wrought
	./wrought

wrought: *.go
	go build .
