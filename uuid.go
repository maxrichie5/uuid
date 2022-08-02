package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/google/uuid"
)

func main() {
	u := uuid.New()
	clipboard.WriteAll(u.String())
	fmt.Println(u.String())
}
