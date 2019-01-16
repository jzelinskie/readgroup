// Package readgroup provides synchronization for reading the same io.Reader
// across N goroutines concurrently.
package readgroup

import (
	"io"
	"io/ioutil"

	"golang.org/x/sync/errgroup"
)

// NewReadGroup specifies the Reader to be consumed.
func NewReadGroup(r io.Reader) *ReadGroup {
	return &ReadGroup{r: r}
}

// A ReadGroup is a collection of goroutines consuming the same io.Reader
// concurrently.
type ReadGroup struct {
	r       io.Reader
	g       errgroup.Group
	writers []io.Writer
}

// Go calls the given function in a new goroutine.
// The first call to return a non-nil error cancels the group; its error will
// be returned by Wait.
func (rg *ReadGroup) Go(f func(io.Reader) error) {
	pr, pw := io.Pipe()
	rg.writers = append(rg.writers, pw)

	rg.g.Go(func() error {
		err := f(pr)
		io.Copy(ioutil.Discard, pr) // Read everything to avoid deadlocks.
		return err
	})
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (rg *ReadGroup) Wait() error {
	rg.g.Go(func() error {
		_, err := io.Copy(io.MultiWriter(rg.writers...), rg.r)
		for _, writer := range rg.writers {
			writer.(io.WriteCloser).Close()
		}
		return err
	})

	return rg.g.Wait()
}
