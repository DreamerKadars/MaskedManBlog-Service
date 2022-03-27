package main

import "MaskedManBlog-Service/MaskedManBlog-Service/route"

func main() {
	route.App.Listen(":8080")
}

