make:
	# sh -c "cd ~/pro/go/src/github.com/PaperMountainStudio/makeweb_gallery; make"
	# sh -c "cd ~/pro/go/src/makeweb; make"
	go vet .
	go test
	go install
	sh -c "cd ~/pro/go/src/github.com/PaperMountainStudio/makeweb-cli; make"
