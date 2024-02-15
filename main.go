package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"

	"log"
)

func main() {

	m := make(map[string]int64)

	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("----------event:", event)
				//if event.Has(fsnotify.Write) {
				//	log.Println("write file:", event.Name)
				//}

				if event.Has(fsnotify.Remove) {
					log.Println("remove file:", event.Name)
				}

				//log.Println("event:", event)
				if event.Has(fsnotify.Rename) {
					log.Println("rename file:", event.Name)
				}

				//update event
				if event.Has(fsnotify.Write) {
					log.Println("write file:", event.Name)
				}

				//create event
				//log.Println("event:", event)
				if event.Has(fsnotify.Create) {
					log.Println("created file:", event.Name)

					ll(event.Name)

				}

				log.Println("========    section 2")

				//===========  get file size
				// Get the fileinfo
				//fileInfo, err := os.Stat("cpnew.txt")

				fileInfo, err := os.Stat(event.Name)
				// Checks for the error
				if err != nil {
					log.Println("=====   temp file issues")
					log.Fatal(err)
				}

				fileSize := fileInfo.Size()
				fmt.Println("Size of the file:", fileSize)

				m[event.Name] = fileSize

				log.Println("===================")
				log.Println(m)

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add("./tmp")
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})

}

func ll(name string) {

}
