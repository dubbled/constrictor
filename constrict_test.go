package constrictor

import (
	"errors"
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

func TestReaderToWriter(t *testing.T) {
	input, err := os.Open("test_input.txt")
	if err != nil {
		t.Error(err)
	}
	defer input.Close()

	var byteCount int64 = 9000
	limitedInput := io.LimitReader(input, byteCount)
	output, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0644)
	if err != nil {
		t.Error(err)
	}

	t1 := time.Now()
	rate := 4000
	written, err := NewReader(limitedInput, rate).WriteTo(output)
	if err != nil {
		t.Error(err)
	}

	checkResults(t, byteCount, written, int64(rate), t1)
}

func TestReaderToBuffer(t *testing.T) {
	input, err := os.Open("test_input.txt")
	if err != nil {
		t.Error(err)
	}
	defer input.Close()

	var byteCount int64 = 4000
	limitedInput := io.LimitReader(input, byteCount)
	output := make([]byte, byteCount)

	t1 := time.Now()
	rate := 2000
	written, err := NewReader(limitedInput, rate).Read(output)
	if err != nil {
		t.Error(err)
	}

	checkResults(t, byteCount, int64(written), int64(rate), t1)
}

func checkResults(t *testing.T, expected, written, rate int64, t1 time.Time) {
	if expected != written {
		t.Error(errors.New("error: input/output mismatch"))
	}

	elapsed := time.Now().Sub(t1).Seconds()
	fmt.Printf("wrote %d bytes in %.02f seconds @ rate %d bytes per second\n", written, elapsed, rate)
}
