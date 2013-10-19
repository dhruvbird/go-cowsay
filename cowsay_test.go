package cowsay

import (
	"fmt"
	"testing"
)

func TestCowsay(t *testing.T) {
	str := Format("Hello    DDD \n\n   World")
	fmt.Printf("%s\n", str)

	str = Format(fmt.Sprint("This is a really long string that ",
		"has a lot of words and should span 2 ",
		"lines."))
	fmt.Printf("%s\n", str)

	str = Format(fmt.Sprint("This is a really long string that ",
		"has a lot more words than the previous line and should ",
		"span more than 2 lines."))
	fmt.Printf("%s\n", str)

	str = Format(fmt.Sprint("ThisLineContainsAReallyLongWord-",
		"supercalifragilisticexpialidocious"))
	fmt.Printf("%s\n", str)

	// Test empty string
	str = Format(fmt.Sprint(""))
	fmt.Printf("%s\n", str)
}
