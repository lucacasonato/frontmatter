package frontmatter

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
)

// ErrMissingDelimiter is the error for when a reader does not contain 2 delimiters
var ErrMissingDelimiter = errors.New("frontmatter: missing delimiter")

// Parse data into frontmatter and rest
func Parse(reader io.Reader, delimiter string) (frontmatter string, rest string, err error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", "", err
	}

	substrings := strings.SplitN(string(data), delimiter, 3)
	if len(substrings) == 3 {
		return strings.TrimSpace(substrings[1]), strings.TrimSpace(substrings[2]), nil
	}

	return "", "", ErrMissingDelimiter
}
