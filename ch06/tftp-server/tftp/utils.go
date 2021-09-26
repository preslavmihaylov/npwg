package tftp

import (
	"bytes"
	"io"
	"strings"
)

func allMust(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

func readString(r *bytes.Buffer, str *string, delim string) error {
	filename, err := r.ReadString(0)
	*str = strings.TrimRight(filename, delim)
	return err
}

func copyN(dst io.Writer, src io.Reader, n int64) error {
	_, err := io.CopyN(dst, src, n)
	if err == io.EOF {
		return nil
	}

	return err
}
