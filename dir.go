package main

import (
	"fmt"
	"log"
	"os"
	"slices"
)
func listDirectories(path string) []string{
	entries, err := os.ReadDir(path)
	list := make([]string,len(entries))
    if err != nil {
        log.Fatal(err)
    }
	for _, e := range entries {
		fmt.Println(e.Name())
		list = append(list,e.Name())
	}
	return list
}
func makeDirectory(path string, foldername string){
	dirPath := path + "\\" + foldername
	fmt.Println(dirPath)
	err := os.Mkdir(dirPath, 0755)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New Directory Created")
}
func changeDirectory(path string, dir string)string{
	list := listDirectories(path)
	if(slices.Contains(list,dir)){
		path = path + "\\" + dir
		fmt.Println(path)
		return path
	}
	log.Fatal("no directory found")
	return path
}

