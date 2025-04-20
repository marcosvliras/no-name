package printer

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdoutPrint(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	StdoutPrint(map[string]string{"hello": "world"})

	w.Close()
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		t.Fatal(err)
	}

	os.Stdout = old

	expected := `{
  "hello": "world"
}
`

	assert.Equal(t, expected, buf.String())
}
