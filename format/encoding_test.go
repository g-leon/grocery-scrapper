package format

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncoding(t *testing.T) {
	testText := "< & >"
	marshaledText, err := json.Marshal(testText)
	if err != nil {
		t.Error(err)
	}

	if bytes.Contains(marshaledText, []byte("\u003c")) {
		t.Errorf("Text does not contain <")
	}
	if bytes.Contains(marshaledText, []byte("\u0026")) {
		t.Errorf("Text does not contain &")
	}
	if bytes.Contains(marshaledText, []byte("\u003e")) {
		t.Errorf("Text does not contain >")
	}

	formatedText := Encoding(marshaledText)
	if !bytes.Contains(formatedText, []byte("\u003c")) {
		t.Errorf("Text does not contain <")
	}
	if !bytes.Contains(formatedText, []byte("\u0026")) {
		t.Errorf("Text does not contain &")
	}
	if !bytes.Contains(formatedText, []byte("\u003e")) {
		t.Errorf("Text does not contain >")
	}
}
