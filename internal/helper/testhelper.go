package helper

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"
	"unsafe"
)

func GetT() *testing.T {
	var buf [8192]byte
	n := runtime.Stack(buf[:], false)
	sc := bufio.NewScanner(bytes.NewReader(buf[:n]))
	for sc.Scan() {
		var p uintptr
		n, _ := fmt.Sscanf(sc.Text(), "testing.tRunner(%v", &p)
		if n != 1 {
			continue
		}
		return (*testing.T)(unsafe.Pointer(p))
	}
	return nil
}
func Fixture(path string) io.Reader {
	log.Println(os.Getwd())
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(b)
}
