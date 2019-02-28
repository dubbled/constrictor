package constrict

import (
	"io"
	"time"
)

type constrictor struct {
	speed   int
	active  bool
	mutable bool
}

type ReadConstrictor struct {
	constrictor
	src io.Reader
}

type WriteConstrictor struct {
	constrictor
	src io.Reader
	dst io.Writer
}

func NewReader(r io.Reader, speed int, mut bool) *ReadConstrictor {
	rc := &ReadConstrictor{
		src: r,
	}
	rc.speed = speed
	return rc
}

func (c *ReadConstrictor) Read() ([]byte, error) {
	c.active = true
	output := []byte{}
	for {
		buf := make([]byte, c.speed)
		n, err := c.src.Read(buf)
		if n > 0 {
			output = append(output, buf[:n]...)
		} else if err == io.EOF {
			break
		} else {
			c.active = false
			return output, err
		}
		time.Sleep(time.Second)
	}

	c.active = false
	return output, nil
}

func NewWriter(src io.Reader, dst io.Writer, speed int, mut bool) *WriteConstrictor {
	wc := &WriteConstrictor{
		src: src,
		dst: dst,
	}

	wc.speed = speed
	wc.mutable = mut

	return wc
}

func (c *WriteConstrictor) Write() error {
	c.active = true
	for {
		buf := make([]byte, c.speed)
		n, err := c.src.Read(buf)
		if n > 0 {
			_, err = c.dst.Write(buf[:n])
			if err != nil {
				return err
			}
		} else if err == io.EOF {
			break
		} else {
			c.active = false
			return err
		}
		time.Sleep(time.Second)
	}
	c.active = false
	return nil
}
