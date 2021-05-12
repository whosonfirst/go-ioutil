package ioutil

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestReadSeekCloserWithFile(t *testing.T) {

	fh, err := os.Open("README.md")

	if err != nil {
		t.Fatalf("Failed to open README.md, %v", err)
	}

	testReadSeekCloser(t, fh)
}

func TestReadSeekCloserWithString(t *testing.T) {

	fh := strings.NewReader("Hello world")
	testReadSeekCloser(t, fh)
}

func testReadSeekCloser(t *testing.T, fh interface{}) {

	rsc, err := NewReadSeekCloser(fh)

	if err != nil {
		t.Fatalf("Failed to create new ReadSeekCloser, %v", err)
	}

	body, err := io.ReadAll(rsc)

	if err != nil {
		t.Fatalf("Failed to read, %v", err)
	}

	_, err = rsc.Seek(0, 0)

	if err != nil {
		t.Fatalf("Failed to seek, %v", err)
	}

	body2, err := io.ReadAll(rsc)

	if err != nil {
		t.Fatalf("Failed to read twice, %v", err)
	}

	if !bytes.Equal(body, body2) {
		t.Fatalf("First and second reads differ")
	}

	err = rsc.Close()

	if err != nil {
		t.Fatalf("Failed to close, %v", err)
	}
}
