package unittest

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" { //jika unit tesnya sama dengan dariwn(mac)
		t.Skip("Unit test tidak bisa jalan karena di mac") //maka akan di skip
	}
	result := HelloWorld("hansen")
	require.Equal(t, "Hello hansen", result)
}
