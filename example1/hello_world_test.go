package example1

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Run main function for before and after tests
func TestMain(m *testing.M) {
	// before run unit testing
	fmt.Println("before run test")

	m.Run()

	// after run unit testing
	fmt.Println("after run test")
}

func TestHelloWithAssert(t *testing.T) {
	result := helloWorld("Mark")
	// Test will assert equal
	assert.Equal(t, "hello mark", result, "result must be 'Hello Mark'")
	// still processed
	fmt.Println("Test with assert done")
}

func TestHelloWithRequire(t *testing.T) {
	result := helloWorld("Mark")
	// Test will require equal
	require.Equal(t, "hello mark", result, "result must be 'Hello Mark'")
	// not processed
	fmt.Println("Test with assert done")
}

func TestHelloWithErrorStop(t *testing.T) {
	result := helloWorld("Mark")

	if result != "Mark" {
		// Test will stop
		t.Error("Result must be Mark")
	}
	// not processed
	fmt.Println("Test Done")
}

func TestHelloWithErrorFail(t *testing.T) {
	result := helloWorld("Mark")

	if result != "Mark" {
		// Test will stop
		t.Fail()
	}
	// will processed
	fmt.Println("Test Done")
}

func TestHelloWithErrorFailNow(t *testing.T) {
	result := helloWorld("Mark")

	if result != "Mark" {
		// Test will stop
		t.FailNow()
	}
	// not processed
	fmt.Println("Test Done")
}

func TestHelloSkip(t *testing.T) {

	if runtime.GOOS == "darwin" {
		// will be skipped but still considered successful
		t.Skip("can't run on MAC OS")
	}
	result := helloWorld("Mark")

	require.Equal(t, "Hello Mark", result, "result must be 'Hello Mark'")
}
