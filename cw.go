package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Assume there are millions of images located in /home/images.
// Write a program to convert each image to grayscale as quickly as possible.

const imagesDir = "/home/images"

var numWorkers = runtime.NumCPU()

var workCh = make(chan string, numWorkers)

func main() {

	paths, err := getImagePaths(imagesDir)

	if err != nil {
		log.Fatal(err)
	}

	go queueWork(paths)

	var w sync.WaitGroup
	w.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, &w)
	}
	w.Wait()
}

func worker(i int, wg *sync.WaitGroup) {

	for path := range workCh {
        convertToGrayscale(path, i)
	}

	wg.Done()
}

func queueWork(paths []string) {
	for _, p := range paths {
		workCh <- p
	}

	close(workCh)
}

// getImagePaths returns all image paths that are in the given directory.
func getImagePaths(directory string) ([]string, error) {
	//imagePaths := []string{}
	//err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
	//	if strings.Contains(path, "png") {
	//		imagePaths = append(imagePaths, path)
	//	}
	//	return nil
	//})
	//if err != nil {
	//	return nil, err
	//}
	s := make([]string, 1e5)
	for i := range s {
		s[i] = "hello" + strconv.Itoa(i)
	}

	return s, nil
}

// convertToGrayscale converts the image at imagePath to grayscale.
func convertToGrayscale(imagePath string, goroutineId int) error {
	//src, err := imaging.Open(imagePath)
	//if err != nil {
	//	return err
	//}

	//img := imaging.Grayscale(src)

	//err = imaging.Save(img, imagePath)
	//if err != nil {
	//	return err
	//}

	fmt.Printf("conveerting greyscale, so much CPU path %s with goroutine %d\n", imagePath, goroutineId)
	time.Sleep(1 * time.Millisecond)

	return nil
}
