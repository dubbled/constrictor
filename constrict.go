package constrictor

import (
	"io"
	"time"
)

type constrictor struct {
	speed int
}

type Reader struct {
	constrictor
	r io.Reader
}

func NewReader(r io.Reader, speed int) *Reader {
	rc := &Reader{
		r: r,
	}
	rc.speed = speed
	return rc
}

func (r *Reader) WriteTo(w io.Writer) (int64, error) {
	throttle := time.Tick(time.Second)
	var count int64
	buf := make([]byte, r.speed)
	for {
		n, err := r.r.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return count, err
		}

		n, err = w.Write(buf[:n])
		if err != nil {
			return count, err
		}

		count += int64(n)
		<-throttle
	}

	return count, nil
}

func (r *Reader) Read(p []byte) (int, error) {
	throttle := time.Tick(time.Second)
	var count int
	buf := make([]byte, r.speed)
	for {
		n, err := r.r.Read(buf)
		if n > 0 {
			p = append(p, buf[:n]...)
			count += n
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return count, err
		}
		<-throttle
	}

	return count, nil
}
