watch:
	clear
	fswatch -o . | xargs -n1 -I{} make cleartest

cleartest:
	clear
	make test

test:
	go test -cover

coverage:
	go test -coverprofile=coverage.out
	go tool cover --html=coverage.out
	rm coverage.out

cover:
	make coverage
