# Go EXIF Data Remover

This Go application checks if an image file contains EXIF metadata and removes it, leaving the image itself intact. This is useful for privacy purposes or if you want to remove unnecessary EXIF information from your photos.

## Features

- Checks for EXIF data in JPEG images.
- Removes EXIF metadata if found.
- Leaves the image intact without modifying the content.

## Requirements

- **Go**: You need to have Go installed. You can download it from [here](https://golang.org/dl/).
- **disintegration/imaging** package: Used for image manipulation.
- **rwcarlsen/goexif** package: Used for reading EXIF data.

### Install the required packages:

```bash
go get -u github.com/disintegration/imaging
go get -u github.com/rwcarlsen/goexif/exif
