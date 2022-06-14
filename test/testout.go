package problems

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func extractStdout(t *testing.T, function func()) string {
	t.Helper()

	orgStdout := os.Stdout
	defer func() {
		os.Stdout = orgStdout
	}()
	r, w, _ := os.Pipe()
	os.Stdout = w

	function()

	w.Close()
	var buffer bytes.Buffer
	if _, err := buffer.ReadFrom(r); err != nil {
		t.Fatalf("failed to read buffer: %v", err)
	}
	// 末尾の改行を除去
	return strings.TrimRight(buffer.String(), "\n")
}
