package metadata_test

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	metadata "github.com/object88/go-image-metadata"
	"github.com/object88/go-image-metadata/jfif"
	"github.com/object88/go-image-metadata/tiff"
)

func Test_Header(t *testing.T) {
	var tcs = []struct {
		name           string
		header         []byte
		expectErr      bool
		expectedReader reflect.Type
	}{
		{"JFIF", []byte{0xff, 0xd8}, false, reflect.TypeOf(&jfif.Reader{})},
		{"Motorola TIFF", []byte{0x4d, 0x4d, 0x00, 0x2a}, false, reflect.TypeOf(&tiff.MotorolaReader{})},
		{"Intel TIFF", []byte{0x49, 0x49, 0x2a, 0x00}, false, reflect.TypeOf(&tiff.IntelReader{})},
		{"bogus Intell TIFF", []byte{0x49, 0x49, 0x01, 0x01}, true, nil},
		{"bogus", []byte{0x01, 0x02}, true, nil},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			buffer := bytes.NewReader(tc.header)
			// buffer := bufio.NewReader(reader)

			r, err := metadata.ReadHeader(buffer)
			if tc.expectErr {
				if err == nil {
					t.Fatal("Expected error reading header; no error returned\n")
				}
			} else {
				if err != nil {
					t.Fatalf("Passed valid header, got err %s", err)
				}
			}

			if tc.expectedReader != nil {
				actualReaderType := reflect.TypeOf(r)
				if r == nil {
					t.Fatal("Expected reader; no reader returned\n")
				} else if actualReaderType != tc.expectedReader {
					t.Fatalf("Expected reader of type %s; got %s", tc.expectedReader, actualReaderType)
				}
			} else {
				if r != nil {
					t.Fatalf("Expected nil reader; got %s\n", reflect.TypeOf(r))
				}
			}
		})
	}
}

func Test_File(t *testing.T) {
	path := "./sample.jpg"
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		t.Fatalf("Failed to open file '%s'; got error %s\n", path, err)
	}
	if file == nil {
		t.Fatalf("Opened file '%s', but did not recieve pointer\n", path)
	}
	defer file.Close()

	// reader := bufio.NewReader(file)
	ir, err := metadata.ReadHeader(file)
	if err != nil {
		t.Fatalf("Error while reading header: %s\n", err)
	}
	ir.Read()
}
