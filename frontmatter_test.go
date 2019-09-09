package frontmatter_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/lucacasonato/frontmatter"
)

func TestParseAll(t *testing.T) {
	const d = "---"
	const f = `name: asd
date: asd`
	const r = `# hello world
hello world`

	fp, rp, err := frontmatter.Parse(strings.NewReader(fmt.Sprintf(`%s
%s
%s
%s`, d, f, d, r)), "---")
	if err != nil {
		t.Fatal(err.Error())
	}
	if fp != f {
		t.Fatalf("frontmatter is %s but should be %s", fp, f)
	}
	if rp != r {
		t.Fatalf("rest is %s but should be %s", rp, r)
	}
}

func TestParseFrontmatterOnly(t *testing.T) {
	const d = "---"
	const f = `name: asd
date: asd`
	fp, rp, err := frontmatter.Parse(strings.NewReader(fmt.Sprintf(`%s
%s
%s`, d, f, d)), "---")
	if err != nil {
		t.Fatal(err.Error())
	}
	if fp != f {
		t.Fatalf("frontmatter is %s but should be %s", fp, f)
	}
	if rp != "" {
		t.Fatalf("rest is %s but should be ''", rp)
	}
}

func TestParseError(t *testing.T) {
	fp, rp, err := frontmatter.Parse(strings.NewReader(""), "---")
	if err == nil {
		t.Fatal("no error on missing delimiter")
	}
	if fp != "" {
		t.Fatalf("frontmatter is %s but should be ''", fp)
	}
	if rp != "" {
		t.Fatalf("rest is %s but should be ''", rp)
	}
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestParseErrorReader(t *testing.T) {
	fp, rp, err := frontmatter.Parse(errReader(0), "---")
	if err == nil {
		t.Fatal("no error on missing delimiter")
	}
	if fp != "" {
		t.Fatalf("frontmatter is %s but should be ''", fp)
	}
	if rp != "" {
		t.Fatalf("rest is %s but should be ''", rp)
	}
}
