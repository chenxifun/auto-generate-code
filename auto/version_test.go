package auto

import (
	"fmt"
	"testing"
)

func TestConvertVersion(t *testing.T) {
	vs := convertVersion(0, "1.0.0.1")
	fmt.Println(vs)

}
