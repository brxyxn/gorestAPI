package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("BASICAPI_DB_USERNAME"), // xisreddd
		os.Getenv("BASICAPI_DB_PASSWORD"), // c3aGVFPauG2BZZUzM5TiuI42dFYW8tSZ
		os.Getenv("BASICAPI_DB_NAME"))     // xisreddd

	a.Run(":3005")
}
