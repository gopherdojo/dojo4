package main

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

type Converter interface {
	convimg() error
}

type jpgConverter struct {
	dist io.Writer
	img  image.Image
}

type pngConverter struct {
	dist io.Writer
	img  image.Image
}

type gifConverter struct {
	dist io.Writer
	img  image.Image
}

func convert(c Converter) error {
	return c.convimg()
}

func (c jpgConverter) convimg() error {
	return jpeg.Encode(c.dist, c.img, nil)
}

func (c pngConverter) convimg() error {
	return png.Encode(c.dist, c.img)
}

func (c gifConverter) convimg() error {
	return gif.Encode(c.dist, c.img, nil)
}
