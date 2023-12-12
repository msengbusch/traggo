package test

import "io"

var _ io.Reader = (*MockReader)(nil)

type MockReader struct {
	Bytes   []byte
	Pointer int
	Err     error
}

func (r *MockReader) Read(p []byte) (n int, err error) {
	if r.Pointer >= len(r.Bytes) {
		if r.Err != nil {
			return 0, r.Err
		}
		return 0, io.EOF
	}

	n = copy(p, r.Bytes[r.Pointer:])
	r.Pointer += n

	return n, nil
}
