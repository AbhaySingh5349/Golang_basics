package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	relativePath := `intermediate\29. File_paths\data\file.txt`
	absolutePath := `E:\Golang_basics\intermediate\29. File_paths\data\file.txt`

	// Join paths using filepath.join
	joinedPath := filepath.Join("home", "Documents", "downloads", "file.zip")
	fmt.Println("Joined Path:", joinedPath)

	normalizedPath := filepath.Clean(`intermediate\29. File_paths\data\file.txt`)
	fmt.Println("Normalized Path:", normalizedPath)

	dir, file := filepath.Split(`E:\Golang_basics\intermediate\29. File_paths\data\file.txt`)
	fmt.Println("File:", file)
	fmt.Println("Dir:", dir)
	fmt.Println("Base Dir: ", filepath.Base(`E:\Golang_basics\intermediate\29. File_paths\data\`))

	fmt.Println("Is relativePath variable absolute:", filepath.IsAbs(relativePath))
	fmt.Println("Is absolutePath variable absolute:", filepath.IsAbs(absolutePath))

	fmt.Println("extension: ", filepath.Ext(file))
	fmt.Println("file_name: ", strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Absolute Path:", absPath)
	}
}
