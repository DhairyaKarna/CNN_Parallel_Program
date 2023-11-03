package scheduler

import (
    "fmt"
    "os"
    "encoding/json"
    "proj1/png"
    "strings"
)

func RunSequential(config Config) {
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
            
            img, err := png.Load(inPath)
            if err != nil{
                panic(err)
            }

            // Apply the specified effects to the image
            for _, effect := range imageEffect.Effects {
                bounds:= img.Bounds
                img.Convolution(effect,bounds.Min.Y, bounds.Max.Y)
            }

            // Save the image to the specified output path
            outPath := fmt.Sprintf("../data/out/%s", id + "_"+ imageEffect.OutPath)
            err = img.Save(outPath)
            if err != nil {
                panic(err)
            }
        }
    }
}
