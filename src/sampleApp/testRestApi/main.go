package main

import restapi "testRestApi/facades"

func main() {
	restapi.Init("api", ":3000")
}
