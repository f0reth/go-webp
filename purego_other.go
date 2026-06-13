//go:build !unix && !darwin && !windows

package webp

import (
	"fmt"
	"image"
	"io"
)

var (
	dynamic    = false
	dynamicErr = fmt.Errorf("webp: platform not supported")
)

func decode(r io.Reader, configOnly, decodeAll bool) (*WEBP, image.Config, error) {
	return nil, image.Config{}, dynamicErr
}

func encode(w io.Writer, m image.Image, quality, method int, lossless, exact bool) error {
	return dynamicErr
}

func loadLibrary(name string) (uintptr, error) {
	return 0, dynamicErr
}
