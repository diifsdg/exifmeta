package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

// checkAndRemoveEXIF checks for EXIF data and removes it
func checkAndRemoveEXIF(filepath string) error {
	// Open the image file
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	// Read EXIF data
	x, err := exif.Decode(file)
	if err != nil {
		// Handle different types of errors
		if err.Error() == "EOF" {
			fmt.Println("No EXIF data found (EOF).")
			return nil
		} else if strings.Contains(err.Error(), "failed to find exif intro marker") {
			fmt.Println("No EXIF data found (EXIF intro marker).")
			return nil
		}
		return fmt.Errorf("error reading EXIF data: %v", err)
	}

	// If EXIF data is present
	if x != nil {
		fmt.Println("EXIF data found, removing it...")

		// Close and reopen the file to remove EXIF
		file.Close()

		// Load the image without EXIF
		img, err := imaging.Open(filepath)
		if err != nil {
			return fmt.Errorf("unable to load image for EXIF removal: %v", err)
		}

		// Save the image, overwriting it without EXIF
		err = imaging.Save(img, filepath)
		if err != nil {
			return fmt.Errorf("unable to save image without EXIF: %v", err)
		}

		fmt.Println("EXIF data successfully removed.")
	} else {
		fmt.Println("No EXIF data found.")
	}

	return nil
}

func main() {
	// Specify the path to the image
	imagePath := "your_image.jpg"

	// Check and remove EXIF
	err := checkAndRemoveEXIF(imagePath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
