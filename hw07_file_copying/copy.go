package main

import (
	"errors"
	"os"
	"io"
	"github.com/cheggaaa/pb/v3"
	"fmt"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	src, err := os.Open(fromPath)
	if err != nil {
		return fmt.Errorf("fromPath - %q: %v", fromPath, err)
	}

	srcStat, err := src.Stat()
	if err != nil {
		return fmt.Errorf("src - %q: %v", fromPath, err)
	}

	if !srcStat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > srcStat.Size() {
		return ErrOffsetExceedsFileSize
	}

	_, err = src.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("src - %q offset - %d: %v", fromPath, offset, err)
	}

	dst, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("toPath - %q: %v", toPath, err)
	}

	l := limit
	if limit == 0 {
		l = srcStat.Size()
	}

	bar := pb.Full.Start64(l)
	defer bar.Finish()
	barReader := bar.NewProxyReader(src)

	if _, err = io.CopyN(dst, barReader, l); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}
	return nil
}
