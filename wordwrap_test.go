package wordwrap

import (
	qt "github.com/frankban/quicktest"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

type testCase struct {
	Name string
	Width uint
}

func getCases() []testCase {
	return []testCase{
		{
			Name: "case01",
			Width: 20,
		},
		{
			Name: "case02",
			Width: 109,
		},
		{
			Name: "case03",
			Width: 109,
		},
	}
}
func TestWrap(t *testing.T) {
	for _, testCase := range getCases() {
		t.Run(testCase.Name, func(t *testing.T){
			c := qt.New(t)
			input := testCase.Name + ".in.txt"
			output := testCase.Name + ".out.txt"

			out := Wrap(getFileContents(input), testCase.Width)
			c.Assert(out, qt.Equals, getFileContents(output))
		})
	}
}


func getFileContents(path string) string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	fileName := basePath + "/testdata/" + path
	if bVal, err := os.ReadFile(fileName); err == nil {
		return string(bVal)
	}else{
		panic(err)
	}
}

