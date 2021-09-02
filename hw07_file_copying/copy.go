package main

import (
	"errors"
	"os"
	"io"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	src, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	srcStat, err := src.Stat()
	if err != nil {
		return err
	}

	if !srcStat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > srcStat.Size() {
		return ErrOffsetExceedsFileSize
	}

	_, err = src.Seek(offset, 0)
	if err != nil {
		return err
	}

	dst, err := os.Create(toPath)
	if err != nil {
		return err
	}

	l := limit
	if limit == 0 {
		l = srcStat.Size()
	}
	if _, err = io.CopyN(dst, src, l); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	return nil
}
