package constrictor

import (
	// "encoding/hex"
	// "errors"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

/* linux only currently
func TestGenerateInput(t *testing.T) {
	input, err := os.Open("/dev/urandom")
	if err != nil {
		t.Error(err)
	}

	var inputLength int64 = 10000
	limitedInput := io.LimitReader(input, inputLength)

	f, err := os.OpenFile("test_input.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Error(err)
	}

	output := hex.NewEncoder(f)
	n, err := io.Copy(output, limitedInput)
	if err != nil {
		t.Error(err)
	} else if n != inputLength {
		errors.New("error: input/output length mismatch")
	}
}
*/

func TestReader(t *testing.T) {
	input, err := os.Open("test_input.txt")
	if err != nil {
		t.Error(err)
	}
	defer input.Close()

	var inputLength int64 = 2000
	limited := io.LimitReader(input, inputLength)

	t1 := time.Now()
	var speed int = 700
	output, err := NewReader(limited, speed).Read()
	if err != nil {
		t.Error(err)
	}

	elapsed := time.Now().Sub(t1).Seconds()
	fmt.Printf("read %d bytes in %.02f seconds\n", len(output), elapsed)
}

func TestWriter(t *testing.T) {
	input, err := os.Open("test_input.txt")
	if err != nil {
		t.Error(err)
	}
	defer input.Close()

	var byteCount int64 = 5000
	limitedInput := io.LimitReader(input, byteCount)
	t1 := time.Now()
	output, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0644)
	if err != nil {
		t.Error(err)
	}

	var speed int = 1000
	if err = NewWriter(output, limitedInput, speed).Write(); err != nil {
		t.Error(err)
	}

	elapsed := time.Now().Sub(t1).Seconds()
	fmt.Printf("wrote %d bytes in %.02f seconds\n", byteCount, elapsed)
}
