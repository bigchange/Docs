package main

import (
	"fmt"
	"sync"
)

func main() {
	var n sync.Map
	n.Store("key", "value")
	fmt.Println(EditDistance("ALGORITHM", "ALTRUISTIC"))
}
