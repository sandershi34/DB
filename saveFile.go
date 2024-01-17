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

func makeDirectory(dir string){
	dirPath := dir
	fmt.Println(dirPath)
	err := os.Mkdir(dirPath,0755)

	if err != nil {
		panic(err)
	}

	fmt.Println(dirPath+" successfully created")
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
func saveFile(filename string,file_contents []byte){
	file,err := os.OpenFile(filename,os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil{
		log.Fatal(err)

	}
	
	defer file.Close()

	_, err = file.Write(file_contents)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success")

}

func makeFolder(foldername string){

}

