package main

import restapi "sampleApp/testRestApi/facades"

func main() {
	restapi.Init("api", "localhost:3000")
}
