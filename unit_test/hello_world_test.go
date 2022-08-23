package unittest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
1. Panic() 			 => akan lansung menghentikan program
2. Fail() 			 => akan mengagalkan unit test, namun tetap melanjutkan eksekusi unit test.
Namun diakhir ketika selesai, maka unit test tersebut dianggap gagal
3. FailNow()		 => akan mengagalkan unit test saat ini juga tanpa melanjutkan eksekusi unit test.
4. Fatal() / Error() => lebih di rekomendasi karena bisa diberi argumen
*/

func TestMain(m *testing.M) {
	fmt.Println("sebelum unit test")

	m.Run() // mengeksekusi semua unit test

	fmt.Println("setelah unit test")
}

func TestHelloWO1(t *testing.T) {
	result := HelloWorld("Hansen")
	if result != "Hello Hansen" {
		// panic("Result is not Hello Hansen")
		t.Fail()
		t.Error("Test 1 gagal")
	}
	// fmt.Println("Test 1 Gagal")
}

func TestHelloWOHan2(t *testing.T) {
	result := HelloWorld("Marcelino")
	if result != "Hello Marcelino" {
		// panic("Result is not Hello Marcelino")
		t.FailNow()
	}

}

func TestHelloWOAzali3(t *testing.T) {
	result := HelloWorld("Azali")
	if result != "Hello Azali" {
		panic("test 3 gagal")
	}
}

func TestSubTest(t *testing.T) {
	t.Run("hansen", func(t *testing.T) {
		result := HelloWorld("Hansen")
		require.Equal(t, "Hello Hansen", result, "result must be `Hello Hansen`")
	})
	t.Run("marcelino", func(t *testing.T) {
		result := HelloWorld("Marcelino")
		require.Equal(t, "Hello Marcelino", result, "result must be `Hello Marcelino`")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hansen",
			request:  "Hansen",
			expected: "Hello Hansen",
		},
		{
			name:     "Marcelino",
			request:  "Marcelino",
			expected: "Hello Marcelino",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

