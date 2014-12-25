package main

import (
	sh "github.com/codeskyblue/go-sh"
)

func main() {
	sh.Command("curl", "-d", "''", "http://localhost:8000/mpesa").Run()
}
