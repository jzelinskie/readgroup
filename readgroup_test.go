package readgroup_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jzelinskie/readgroup"
)

// SimpleUsage illustrates a very basic use case where you need to make
// transformations to the same io.Reader, but it's too large for working
// memory.
//
// Instead of a large streaming file, a small string is used.
func ExampleReadGroup_simpleUsage() {
	buf := bytes.NewBufferString("My Huge Streaming Video Reader")

	rg := readgroup.NewReadGroup(buf)

	rg.Go(func(r io.Reader) error {
		_, err := io.Copy(os.Stdout, r)
		io.WriteString(os.Stdout, "\n")
		return err
	})

	rg.Go(func(r io.Reader) error {
		hellostr, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}

		fmt.Println(strings.ToLower(string(hellostr)))
		return nil
	})

	rg.Go(func(r io.Reader) error {
		return errors.New("error that could short-circuit the go routines")
	})

	err := rg.Wait()
	if err != nil {
		fmt.Printf("\n%s", err.Error())
	}
}
