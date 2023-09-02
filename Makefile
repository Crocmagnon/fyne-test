run:
	go run main.go

build: fyne-test
	fyne build

clean:
	rm fyne-test

build-linux:
	fyne-cross linux -arch='*'

build-windows:
	fyne-cross windows -arch='*'
