package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Farul")

	if result != "Hello Farul" {
		//error
		panic("Result is not 'Hello Farul'")
	}
}

func TestHelloWorldClone(t *testing.T) {
	result := HelloWorld("Farul Wananda")

	if result != "Hello Farul Wananda" {
		//error
		panic("Result is not 'Hello Farul Wananda'")
	}
}