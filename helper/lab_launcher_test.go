package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Farul", func(t *testing.T) {
		result := HelloWorld("Farul")
		require.Equal(t, "Hello Farul", result, "Result must be 'Hello Farul'")
	})
	t.Run("Wananda", func(t *testing.T) {
		result := HelloWorld("Wananda")
		require.Equal(t, "Hello Wananda", result, "Result must be 'Hello Wananda'")
	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Test cannot run on Windows")
	}

	result := HelloWorld("Farul")
	require.Equal(t, "Hello Farul", result, "Result mus be 'Hello Farul'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Farul")
	assert.Equal(t, "Hello Farul", result, "Result mus be 'Hello Farul'")
	fmt.Println("TestHelloWorldAssert is DONE")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Farul")
	require.Equal(t, "Hello Farul", result, "Result mus be 'Hello Farul'")
	fmt.Println("TestHelloWorldRequire is DONE")
}

func TestHelloWorldOrigin(t *testing.T) {
	result := HelloWorld("Farul")

	if result != "Hello Farul" {
		//error
		//* t.Fail()
		t.Error("Result must be 'Hello Farul'")
	}

	fmt.Println("TestHelloWorldOrigin is DONE")
}

func TestHelloWorldClone(t *testing.T) {
	result := HelloWorld("Farul Wananda")

	if result != "Hello Farul Wananda" {
		//error
		//* t.FailNow()
		t.Fatal("Result must be 'Hello Farul Wananda'")
	}

	fmt.Println("TestHelloWorldClone is DONE")
}
