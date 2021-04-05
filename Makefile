#all: browser.out 
all: browser.out browser.exe

browser.out : browser.go
	echo "Building for native"
	go build -o browser.out browser.go
	
browser.exe: browser.go
	echo "Building for windows"
	GOOS=windows go build browser.go
	cp browser.exe /mnt/webdav

#arg.out : arg.go
#	echo "Building for native"
#	go build -o arg.out arg.go
#	cp arg.out /mnt/webdav
#	
#arg.exe: arg.go
#	echo "Building for windows"
#	GOOS=windows go build arg.go
#	cp arg.exe /mnt/webdav

clean:
	rm -f browser.out browser.exe arg.out arg.exe
