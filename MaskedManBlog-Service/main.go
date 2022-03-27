package main

import "github.com/DreamerKadars/MaskedManBlog-Service/MaskedManBlog-Service/route"
func main() {
	route.App.Listen(":8080")
}

