// Package png allows for loading png images and applying
// image flitering effects on them.
package png

import (
	"image"
	"image/color"
)

// Grayscale applies a grayscale filtering effect to the image
func (img *Image) Grayscale() {

	// Bounds returns defines the dimensions of the image. Always
	// use the bounds Min and Max fields to get out the width
	// and height for the image
	bounds := img.out.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			//Returns the pixel (i.e., RGBA) value at a (x,y) position
			// Note: These get returned as int32 so based on the math you'll
			// be performing you'll need to do a conversion to float64(..)
			r, g, b, a := img.in.At(x, y).RGBA()

			//Note: The values for r,g,b,a for this assignment will range between [0, 65535].
			//For certain computations (i.e., convolution) the values might fall outside this
			// range so you need to clamp them between those values.
			greyC := clamp(float64(r+g+b) / 3)

			//Note: The values need to be stored back as uint16 (I know weird..but there's valid reasons
			// for this that I won't get into right now).
			img.out.Set(x, y, color.RGBA64{greyC, greyC, greyC, uint16(a)})
		}
	}
}


// Sharpen applies a sharpen effect to the image
func (img *Image) Sharpen() {
    // The sharpen kernel as a 3x3 matrix
    kernel := [3][3]float64{
        {0, -1, 0},
        {-1, 5, -1},
        {0, -1, 0},
    }

    // Get image bounds
    bounds := img.out.Bounds()

    // Create a temporary image to store the result
    tempImg := image.NewRGBA64(bounds)

    for y := bounds.Min.Y + 1; y < bounds.Max.Y - 1; y++ {
        for x := bounds.Min.X + 1; x < bounds.Max.X - 1; x++ {
            var sumR, sumG, sumB float64
            // Apply the kernel to each color channel
            for ky := -1; ky <= 1; ky++ {
                for kx := -1; kx <= 1; kx++ {
                    r, g, b, _ := img.in.At(x+kx, y+ky).RGBA()
                    kernelValue := kernel[ky+1][kx+1]
                    sumR += float64(r) * kernelValue
                    sumG += float64(g) * kernelValue
                    sumB += float64(b) * kernelValue
                }
            }
            // Clamp the results to [0, 65535]
            newR := clamp(sumR)
            newG := clamp(sumG)
            newB := clamp(sumB)

            // Set the new color value to the temporary image
            tempImg.Set(x, y, color.RGBA64{newR, newG, newB, 65535})
        }
    }

    // Copy the temporary image to the output image
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, a := tempImg.At(x, y).RGBA()
            img.out.Set(x, y, color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)})
        }
    }
}

// EdgeDetection applies an edge detection effect to the image
func (img *Image) EdgeDetection() {
    // The edge detection kernel as a 3x3 matrix
    kernel := [3][3]float64{
        {-1, -1, -1},
        {-1, 8, -1},
        {-1, -1, -1},
    }

    // Get image bounds
    bounds := img.out.Bounds()

    // Create a temporary image to store the result
    tempImg := image.NewRGBA64(bounds)

    for y := bounds.Min.Y + 1; y < bounds.Max.Y - 1; y++ {
        for x := bounds.Min.X + 1; x < bounds.Max.X - 1; x++ {
            var sumR, sumG, sumB float64
            // Apply the kernel to each color channel
            for ky := -1; ky <= 1; ky++ {
                for kx := -1; kx <= 1; kx++ {
                    r, g, b, _ := img.in.At(x+kx, y+ky).RGBA()
                    kernelValue := kernel[ky+1][kx+1]
                    sumR += float64(r) * kernelValue
                    sumG += float64(g) * kernelValue
                    sumB += float64(b) * kernelValue
                }
            }
            // Clamp the results to [0, 65535]
            newR := clamp(sumR)
            newG := clamp(sumG)
            newB := clamp(sumB)

            // Set the new color value to the temporary image
            tempImg.Set(x, y, color.RGBA64{newR, newG, newB, 65535})
        }
    }

    // Copy the temporary image to the output image
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, a := tempImg.At(x, y).RGBA()
            img.out.Set(x, y, color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)})
        }
    }
}

// Blur applies a blur effect to the image
func (img *Image) Blur() {
    // The blur kernel as a 3x3 matrix
    kernel := [3][3]float64{
        {1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
        {1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
        {1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
    }

    // Get image bounds
    bounds := img.out.Bounds()

    // Create a temporary image to store the result
    tempImg := image.NewRGBA64(bounds)

    for y := bounds.Min.Y + 1; y < bounds.Max.Y - 1; y++ {
        for x := bounds.Min.X + 1; x < bounds.Max.X - 1; x++ {
            var sumR, sumG, sumB float64
            // Apply the kernel to each color channel
            for ky := -1; ky <= 1; ky++ {
                for kx := -1; kx <= 1; kx++ {
                    r, g, b, _ := img.in.At(x+kx, y+ky).RGBA()
                    kernelValue := kernel[ky+1][kx+1]
                    sumR += float64(r) * kernelValue
                    sumG += float64(g) * kernelValue
                    sumB += float64(b) * kernelValue
                }
            }
            // Clamp the results to [0, 65535]
            newR := clamp(sumR)
            newG := clamp(sumG)
            newB := clamp(sumB)

            // Set the new color value to the temporary image
            tempImg.Set(x, y, color.RGBA64{newR, newG, newB, 65535})
        }
    }

    // Copy the temporary image to the output image
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, a := tempImg.At(x, y).RGBA()
            img.out.Set(x, y, color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)})
        }
    }
}
