package scheduler

import (
	"encoding/json"
	"os"
	"strings"
	"proj1/png"
	"fmt"
	"sync"
)

func processSlice(img *png.Image, startY, endY int, effect string, wg *sync.WaitGroup) {
	// Ensure the WaitGroup counter is decremented when this function returns
	defer wg.Done()

	// Apply the specified effect to the image
	img.Convolution(effect, startY, endY)
	
}

func RunParallelSlices(config Config) {
	// Load the JSON strings from the effects.txt file
    effectsPathFile := fmt.Sprintf("../data/effects.txt")
    effectsFile, _ := os.Open(effectsPathFile)
    reader := json.NewDecoder(effectsFile)

    dataDirs := config.DataDirs
    identifier := strings.Split(dataDirs, "+")

	for {

		var imageEffect ImageEffect
		err := reader.Decode(&imageEffect)

		// Stop at EOF.
		if err != nil {
			break
		}

		for _, id := range identifier {
			

			// Load the image from the specified input path
			inPath := fmt.Sprintf("../data/in/%s/%s", id, imageEffect.InPath)
			// Save the image to the specified output path
			outPath := fmt.Sprintf("../data/out/%s", id + "_"+ imageEffect.OutPath)

			img, err := png.Load(inPath)
			if err != nil{
				panic(err)
			}

			// Determine the number of rows to process per Goroutine
			rowsPerGoroutine := img.Bounds.Dy() / config.ThreadCount

			// Apply the specified effects to the image
			for _, effect := range imageEffect.Effects {
				// Create a WaitGroup to wait for all Goroutines to complete
				var wg sync.WaitGroup

				// Launch a Goroutine for each slice of the image
				for i := 0; i < config.ThreadCount; i++ {
					startY := i * rowsPerGoroutine
					endY := startY + rowsPerGoroutine
					if i == config.ThreadCount-1 {
						endY = img.Bounds.Dy()  // Ensure the last Goroutine processes all remaining rows
					}

					wg.Add(1)
					go processSlice(img, startY, endY, effect, &wg)
				}

				// Wait for all Goroutines to complete
				wg.Wait()
			}

			err = img.Save(outPath)
			if err != nil {
				panic(err)
			}
		}
	}
}
