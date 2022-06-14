package gui

import (
	"fmt"
	"testing"
)

func TestFindRegionId(*testing.T) {
	for _, r := range regionData {
		fmt.Println(findRegionId(r))
	}
}
