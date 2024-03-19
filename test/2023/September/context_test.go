package September

import (
	"fmt"
	"os"
	"testing"
)

func TestContxtPtr(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err = %v\n", err)
		return
	}
	fmt.Printf("dir = %v\n", dir)
}
