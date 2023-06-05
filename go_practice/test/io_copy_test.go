// Copyright © Bytedance All Rights Reserved.
// :author Alex Ma
// :date 2022/10/28

package test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func ReadAll() []byte {
	dir := "/Users/machao/Downloads/00000053a43d11ebb49fac1f6b0eb9d0"
	f, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return data
}

func IOCopy() []byte {
	dir := "/Users/machao/Downloads/00000053a43d11ebb49fac1f6b0eb9d0"
	f, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	data := bytes.NewBuffer(nil)
	_, err = io.Copy(data, f)
	if err != nil {
		panic(err)
	}
	return data.Bytes()
}

func IOCopyN() []byte {
	dir := "/Users/machao/Downloads/00000053a43d11ebb49fac1f6b0eb9d0"
	f, err := os.Open(dir)
	if err != nil {
		panic(err)
	}

	data := bytes.NewBuffer(nil)
	for {
		_, err := io.CopyN(data, f, 1024)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
	}
	return data.Bytes()
}

func TestIO(t *testing.T) {
	fmt.Println(len(ReadAll()))
	//fmt.Println(len(IOCopy()))
	fmt.Println(len(IOCopyN()))
	fmt.Println(string(IOCopyN()) == string(ReadAll()))
}

func BenchmarkIOReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadAll()
	}
}

func BenchmarkIOCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IOCopy()
	}
}
