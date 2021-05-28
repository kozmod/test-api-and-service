package main

import (
	"fmt"
	"test-api-and-service/api/meta"
)

func main() {
	fmt.Printf(meta.Some{}.String())
}
