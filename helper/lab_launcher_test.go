package helper

import (
	"fmt"
	"testing"
)

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