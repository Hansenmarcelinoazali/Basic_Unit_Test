package unittest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("hansen")
	require.Equal(t, "Hello hansen",result,"must be hello hansen")
	fmt.Println("tidak di eksekusi")
}