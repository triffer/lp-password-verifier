package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

import . "github.com/triffer/lp-password-verifier/internal"

func TestSmokeTest(t *testing.T) {

	file, _ := filepath.Abs("testdata/export.csv")
	os.Args = []string{"main", file}

	out := captureOutput(func() {
		Main()
	})

	assert.Equal(t, "The following passwords have been pwned\n\nName: Number1 \nPassword: 123456789 \n--------------------------------\n", out)
}

func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	err := w.Close()

	if err != nil {
		log.Fatalln(err)
	}

	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}
