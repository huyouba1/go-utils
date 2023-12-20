package go_utils

import (
	"bytes"
        "fmt"
	"compress/zlib"
	"io"
)

// 压缩，zlib deflate压缩算法
func Compress(in []byte) (result *bytes.Buffer, err error) {
	result = new(bytes.Buffer)
	w := zlib.NewWriter(result)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			w.Close()
		}
	}()
	_, err = w.Write(in)
	if err != nil {
		fmt.Println("During compression, buffer write exception")
		return nil, err
	}
	if err = w.Close(); err != nil {
		return nil, err
	}
	return
}

// 数据解压缩，zlib deflate压缩算法
func DeCompress(in []byte) (result *bytes.Buffer, err error) {
	result = new(bytes.Buffer)
	r, err := zlib.NewReader(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	_, err = io.Copy(result, r)
	if err != nil {
		return
	}
	return
}

