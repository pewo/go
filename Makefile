#all: browser.out 
all: browser.out browser.exe tlclient.exe tlclient.out

browser.out : browser.go
	echo "Building for native"
	go build -o browser.out browser.go
	
browser.exe: browser.go
	echo "Building for windows"
	GOOS=windows go build browser.go
	cp browser.exe /mnt/webdav

tlclient.out : tlclient.go
	echo "Building for native"
	go build -o tlclient.out tlclient.go
	cp tlclient.out /mnt/webdav
	
tlclient.exe: tlclient.go
	echo "Building for windows"
	GOOS=windows go build tlclient.go
	cp tlclient.exe /mnt/webdav

clean:
	rm -f browser.out browser.exe tlclient.out tlclient.exe
