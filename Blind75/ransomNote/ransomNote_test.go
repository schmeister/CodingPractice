package ransomNote

import (
	"fmt"
	"testing"
)

func TestCanConstruct(t *testing.T) {

	fmt.Println(CanConstruct("This is a test", "more words"))
	fmt.Println(CanConstruct("This is a test", "Tthis is a tes"))

}
