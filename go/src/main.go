package main

import "hi/greeting"

// Couple of points
// 1. Created a go.mod file; appears to be a module of sort
// 2. Changed the import so that it looks for the package relatively to the module name
// 3. I am not totally sure how GOPATH works in this case
// 4. Most useful resource = https://stackoverflow.com/questions/70314880/golang-looks-for-packages-in-goroot-instead-of-gopath

func main() {
	greeting.Hello()
	greeting.Hi()
}
