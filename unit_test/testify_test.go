package unittest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

)

func TestHelloWorldAssertion(t *testing.T) {
	fmt.Println("ini assertion")

	result := HelloWorld("Hansen")
	assert.Equal(t, "Hello Hansen", result, "Result must be Hello Hansen") //hasil yang diharapkan
	fmt.Println("Dieksekusi")
}
