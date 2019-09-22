package replaceablewriter

import (
	"io"
	"sync"
)

type Writer struct {
	w  io.Writer
	mu sync.Mutex
}

// New returns new replaceablewriter
func New(w io.Writer) *Writer {
	return &Writer{w: w}
}

// Write implements io.Writer
func (w *Writer) Write(b []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.w.Write(b)
}

// Close implements io.Closer
func (w *Writer) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if closer, ok := w.w.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

// Replace internal io.Writer
func (w *Writer) Replace(wtr io.Writer) {
	w.mu.Lock()
	if closer, ok := w.w.(io.Closer); ok {
		defer closer.Close()
	}
	defer w.mu.Unlock()
	w.w = wtr
}
