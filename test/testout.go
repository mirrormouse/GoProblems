package problems

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

//スライス同士の比較

func compslice(s, t []int) bool {
	size := len(s)
	if size == len(t) {
		for i := 0; i < size; i++ {
			if s[i] != t[i] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

//スライスの文字列化
func slicetostring(s []int) string {
	t := "{"
	for i := 0; i < len(s); i++ {
		if i > 0 {
			t += ","
		}
		t += strconv.Itoa(s[i])

	}
	t += "}"
	return t
}

//スライスの文字列化
func slicestrtostring(s []string) string {
	t := "{"
	for i := 0; i < len(s); i++ {
		if i > 0 {
			t += ","
		}
		t += s[i]

	}
	t += "}"
	return t
}

//2次元スライスの文字列化
func slice2strtostring(s [][]string) string {
	t := "{"
	for i := 0; i < len(s); i++ {
		t += "{"
		for j := 0; j < len(s[i]); j++ {
			if j > 0 {
				t += ","
			}
			t += s[i][j]

		}
		t += "},"
	}
	t += "}"
	return t
}

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
