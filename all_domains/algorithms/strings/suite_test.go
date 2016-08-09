package strings

// File contains hook to link gocheck with Golang test system
import (
	"bytes"
	. "gopkg.in/check.v1"
	"io"
	"os"
	"testing"
)

func captureStdout(f func(), str ...string) string {
	oldStdout := os.Stdout
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	for _, s := range str {
		os.Stdin.Read([]byte(s))
	}
	os.Stdout = w
	os.Stdin = r

	f()

	w.Close()
	os.Stdout = oldStdout
	os.Stdin = oldStdin
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// Hook up gocheck into the "go test" runner.
func TestStart(t *testing.T) {
	TestingT(t)
}
