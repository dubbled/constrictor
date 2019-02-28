package constrict

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func TestReader(t *testing.T) {
	input, err := os.Open("/dev/urandom")
	if err != nil {
		t.Error(err)
	}
	defer input.Close()

	limited := io.LimitReader(input, 2000)
	t1 := time.Now()
	output, err := NewReader(limited, 700, false).Read()
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	fmt.Printf("read %d bytes in %f seconds\n", len(output), time.Now().Sub(t1).Seconds())
}

func TestWriter(t *testing.T) {
	input, err := os.Open("/dev/urandom")
	if err != nil {
		t.Error(err)
	}
	defer input.Close()

	var byteCount int64
	byteCount = 5000
	limited := io.LimitReader(input, byteCount)
	t1 := time.Now()
	output, err := os.OpenFile("/dev/null", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	if err = NewWriter(limited, output, 300, false).Write(); err != nil {
		t.Error(err)
	}

	fmt.Printf("wrote %d bytes in %f seconds\n", byteCount, time.Now().Sub(t1).Seconds())
}
