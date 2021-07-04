package main

import "go-tests/api"

func main() {
	api.GetRouter().Run("0.0.0.0:8000")
}
