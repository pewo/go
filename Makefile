#all: browser.out 
all: browser.out browser.exe

browser.out : browser.go
	echo "Building for native"
	go build -o browser.out browser.go
	./browser.out

browser.exe: browser.go
	echo "Building for windows"
	GOOS=windows go build browser.go
	cp browser.exe /mnt/webdav

clean:
	rm -f browser.out browser.exe
