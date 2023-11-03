package png

import(
	"image"
	"image/color"
)

func (img *Image) Convolution(effect string, startY, endY int){
	switch effect {
	case "S":
		kernel := []float64{0, -1, 0, -1, 5, -1, 0, -1, 0}
		img.Convolve(kernel, startY, endY)
	case "E":
		kernel := []float64{-1, -1, -1, -1, 8, -1, -1, -1, -1}
		img.Convolve(kernel, startY, endY)
	case "B":
		kernel := []float64{1 / 9.0, 1 / 9.0, 1 / 9.0, 1 / 9.0, 1 / 9.0, 1 / 9.0, 1 / 9.0, 1 / 9.0, 1 / 9.0}
		img.Convolve(kernel, startY, endY)
	case "G":
		img.ApplyGrayscale(startY, endY)
	default:
		panic("Unknown effect: " + effect)
	}
}


// Convolve applies the specified kernel to a slice of the image
// between startY and endY.
func (img *Image) Convolve(kernel []float64, startY int, endY int) {
	bounds := img.out.Bounds()
	// Create a temporary image to hold the intermediate results
	tempImg := image.NewRGBA64(bounds)

	for y := startY; y < endY; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var sumR, sumG, sumB float64
			// Iterate over the 3x3 kernel grid
			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					// Get the color of the current pixel
					r, g, b, _ := img.in.At(x+kx, y+ky).RGBA()
					// Get the corresponding kernel value
					kernelValue := kernel[(ky+1)*3+(kx+1)]
					// Accumulate the sum for each color channel
					sumR += float64(r) * kernelValue
					sumG += float64(g) * kernelValue
					sumB += float64(b) * kernelValue
				}
			}
			// Set the color of the current pixel in the temporary image
			tempImg.Set(x, y, color.RGBA64{
				uint16(clamp(sumR)),
				uint16(clamp(sumG)),
				uint16(clamp(sumB)),
				uint16(65535),  // Assume full alpha
			})
		}
	}

	// Copy the temporary image back to the output image
	for y := startY; y < endY; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			img.out.Set(x, y, tempImg.At(x, y))
		}
	}
}


func (img *Image) ApplyGrayscale(startY, endY int) {
	bounds := img.out.Bounds()
	for y := startY; y < endY; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.in.At(x, y).RGBA()
			greyC := clamp(float64(r+g+b) / 3.0)
			img.out.Set(x, y, color.RGBA64{greyC, greyC, greyC, uint16(a)})
		}
	}
}