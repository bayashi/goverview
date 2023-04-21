package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/bayashi/actually"
	pt "github.com/bayashi/go-proptree"
)

func gomodFile() string {
	return `module github.com/bayashi/goverviewtest

go 1.19
`
}

func mainFile() string {
	return `package main

const X = "x"

func main() {
	println("main")
}

func X() {
	println(X)
}
`
}

func licenseFile() string {
	return `MIT License

Copyright (c) 2023 Dai Okabayashi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`
}

// general
func expect1(dir string) string {
	return "\n" +
		"┌ " + dir + "/\n" +
		"├── .gitignore\n" +
		"├── LICENSE: License MIT\n" +
		"├── bar.txt\n" +
		"├── go.mod: go 1.19\n" +
		"└──* main.go: main\n" +
		"      Func: X\n" +
		"      Const: X\n"
}

// With private func
func expect2(dir string) string {
	return "\n" +
		"┌ " + dir + "/\n" +
		"├── .gitignore\n" +
		"├── LICENSE: License MIT\n" +
		"├── bar.txt\n" +
		"├── go.mod: go 1.19\n" +
		"└──* main.go: main\n" +
		"      Func: X\n" +
		"      func: main\n" +
		"      Const: X\n"
}

// Without bar.txt
func expect3(dir string) string {
	return "\n" +
		"┌ " + dir + "/\n" +
		"├── .gitignore\n" +
		"├── LICENSE: License MIT\n" +
		"├── go.mod: go 1.19\n" +
		"└──* main.go: main\n" +
		"      Func: X\n" +
		"      Const: X\n"
}

func createFile(path string, content string) {
	f, _ := os.Create(path)
	defer f.Close()
	_, err := f.Write([]byte(content))
	if err != nil {
		panic("could not write: " + path + ", " + content)
	}
}

func TestGoProject(t *testing.T) {
	temp, _ := filepath.Abs(t.TempDir())
	base := filepath.Base(temp)
	//fmt.Printf("%#v\n", temp)
	createFile(filepath.Join(temp, "LICENSE"), licenseFile())
	createFile(filepath.Join(temp, "go.mod"), gomodFile())
	createFile(filepath.Join(temp, "main.go"), mainFile())
	createFile(filepath.Join(temp, ".gitignore"), "foo.txt")
	createFile(filepath.Join(temp, "foo.txt"), "")
	createFile(filepath.Join(temp, "bar.txt"), "")

	tts := []struct {
		o      *options
		expect string
	}{
		{o: &options{path: temp}, expect: expect1(base)},
		{o: &options{path: temp, showAll: true}, expect: expect2(base)},
		{o: &options{path: temp, ignore: []string{"bar.txt"}}, expect: expect3(base)},
	}
	for _, tt := range tts {
		tree, err := fromLocal(tt.o)
		actually.Got(err).Nil(t)
		buf := &bytes.Buffer{}
		tree.RenderAsText(buf, pt.RenderTextDefaultOptions())
		actually.Got(buf.String()).Expect(tt.expect).ShowRawData().Same(t)
	}
}
