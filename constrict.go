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
		read, err := r.r.Read(buf)
		if read > 0 {
			written, err := w.Write(buf[:read])
			count += int64(written)
			if err != nil {
				return count, err
			}
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

func (r *Reader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	throttle := time.Tick(time.Second)
	var count int
	buf := make([]byte, r.speed)
	for {
		read, err := r.r.Read(buf)
		if read > 0 {
			p = append(p, buf[:read]...)
			count += read
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
