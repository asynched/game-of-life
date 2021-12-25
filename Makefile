GC := go
SOURCE := main.go

build:
	cd src &&\
	go build -o main.out main.go &&\
	mv main.out ..
