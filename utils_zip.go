package wxxx

import (
	"bytes"
	"compress/zlib"
	"io"
)

func deflateZip(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	if _, err := w.Write(src); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func deflateUnZip(src []byte) ([]byte, error) {
	buf := bytes.NewReader(src)
	var out bytes.Buffer
	r, err := zlib.NewReader(buf)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(&out, r)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
