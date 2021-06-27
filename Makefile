main:
	go build $(WINDOWS_GUI) -o=.\build\JcopWebtoonDownloader.exe .\src
	cd build && .\JcopWebtoonDownloader.exe

nodebug: WINDOWS_GUI = -ldflags -H=windowsgui

nodebug: main

testlib:
	go test .\src\WTdown

gitpub:
	set /p comment=Type Comment: 
	git init
	git add .
	git commit -m "%comment%"
	git push -f origin master

