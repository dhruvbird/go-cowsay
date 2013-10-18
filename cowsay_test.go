package cowsay

import (
       "fmt"
       "testing"
)

func TestCowsay(t *testing.T) {
     str := format("Hello World")
     fmt.Printf("%s\n", str)
}
