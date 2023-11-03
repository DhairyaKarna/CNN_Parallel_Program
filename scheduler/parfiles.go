package scheduler

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"proj1/png"
	"sync"
)

func processImage(imageEffect ImageEffect, id string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Load the image from the specified input path
	inPath := fmt.Sprintf("../data/in/%s/%s", id, imageEffect.InPath)
	// Save the image to the specified output path
	outPath := fmt.Sprintf("../data/out/%s", id + "_"+ imageEffect.OutPath)

	img, err := png.Load(inPath)
	if err != nil{
		panic(err)
	}

	// Apply the specified effects to the image
	for _, effect := range imageEffect.Effects {
		bounds:= img.Bounds
		img.Convolution(effect,bounds.Min.Y, bounds.Max.Y)
	}

	// Save the processed image to the specified output path
	err = img.Save(outPath)
	if err != nil {
		panic(err)
	}
}

func RunParallelFiles(config Config) {
	var wg sync.WaitGroup

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
			wg.Add(1)
			go processImage(imageEffect, id, &wg)
		}
		wg.Wait()
	}
}
