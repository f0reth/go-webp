package webp

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"testing"
)

//go:embed testdata/test.webp
var testWebp []byte

//go:embed testdata/test.png
var testPng []byte

//go:embed testdata/anim.webp
var testWebpAnim []byte

func skipIfNoLibrary(tb testing.TB) {
	if err := Dynamic(); err != nil {
		fmt.Println(err)
		tb.Skip()
	}
}

func TestDecode(t *testing.T) {
	skipIfNoLibrary(t)

	img, _, err := decode(bytes.NewReader(testWebp), false, false)
	if err != nil {
		t.Fatal(err)
	}

	w, err := writeCloser()
	if err != nil {
		t.Fatal(err)
	}

	err = jpeg.Encode(w, img.Image[0], nil)
	if err != nil {
		t.Error(err)
	}
}

func TestDecodeAnim(t *testing.T) {
	skipIfNoLibrary(t)

	ret, _, err := decode(bytes.NewReader(testWebpAnim), false, true)
	if err != nil {
		t.Fatal(err)
	}

	if len(ret.Image) != len(ret.Delay) {
		t.Errorf("got %d, want %d", len(ret.Delay), len(ret.Image))
	}

	if len(ret.Image) != 17 {
		t.Errorf("got %d, want %d", len(ret.Image), 17)
	}

	for _, img := range ret.Image {
		w, err := writeCloser()
		if err != nil {
			t.Fatal(err)
		}

		err = jpeg.Encode(w, img, nil)
		if err != nil {
			t.Error(err)
		}

		err = w.Close()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestImageDecode(t *testing.T) {
	skipIfNoLibrary(t)

	img, _, err := image.Decode(bytes.NewReader(testWebp))
	if err != nil {
		t.Fatal(err)
	}

	w, err := writeCloser()
	if err != nil {
		t.Fatal(err)
	}

	err = jpeg.Encode(w, img, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestImageDecodeAnim(t *testing.T) {
	skipIfNoLibrary(t)

	img, _, err := image.Decode(bytes.NewReader(testWebpAnim))
	if err != nil {
		t.Fatal(err)
	}

	w, err := writeCloser()
	if err != nil {
		t.Fatal(err)
	}

	err = jpeg.Encode(w, img, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestDecodeConfig(t *testing.T) {
	skipIfNoLibrary(t)

	cfg, err := DecodeConfig(bytes.NewReader(testWebp))
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Width != 512 {
		t.Errorf("width: got %d, want %d", cfg.Width, 512)
	}

	if cfg.Height != 512 {
		t.Errorf("height: got %d, want %d", cfg.Height, 512)
	}
}

func TestImageDecodeConfig(t *testing.T) {
	skipIfNoLibrary(t)

	cfg, _, err := image.DecodeConfig(bytes.NewReader(testWebp))
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Width != 512 {
		t.Errorf("width: got %d, want %d", cfg.Width, 512)
	}

	if cfg.Height != 512 {
		t.Errorf("height: got %d, want %d", cfg.Height, 512)
	}
}

func TestEncodeRGBA(t *testing.T) {
	skipIfNoLibrary(t)

	img, err := png.Decode(bytes.NewReader(testPng))
	if err != nil {
		t.Fatal(err)
	}

	w, err := writeCloser()
	if err != nil {
		t.Fatal(err)
	}

	err = Encode(w, img)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncode(t *testing.T) {
	skipIfNoLibrary(t)

	img, err := Decode(bytes.NewReader(testWebp))
	if err != nil {
		t.Fatal(err)
	}

	w, err := writeCloser()
	if err != nil {
		t.Fatal(err)
	}

	err = encode(w, img, DefaultQuality, DefaultMethod, false, false)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkDecode(b *testing.B) {
	skipIfNoLibrary(b)

	for b.Loop() {
		_, _, err := decode(bytes.NewReader(testWebp), false, false)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	skipIfNoLibrary(b)

	img, err := Decode(bytes.NewReader(testWebp))
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		err = encode(io.Discard, img, DefaultQuality, DefaultMethod, false, false)
		if err != nil {
			b.Error(err)
		}
	}
}

type discard struct{}

func (d discard) Close() error {
	return nil
}

func (discard) Write(p []byte) (int, error) {
	return len(p), nil
}

var discardCloser io.WriteCloser = discard{}

func writeCloser(s ...string) (io.WriteCloser, error) {
	if len(s) > 0 {
		f, err := os.Create(s[0])
		if err != nil {
			return nil, err
		}

		return f, nil
	}

	return discardCloser, nil
}
