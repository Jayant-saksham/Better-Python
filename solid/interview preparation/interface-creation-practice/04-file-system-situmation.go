package interface_creation_practice

import "fmt"

type FileSystemEntry interface {
	Size() int
}

type File struct {
	Name    string
	Content string
}
type Directory struct {
	Name  string
	files []FileSystemEntry
}

func (f File) Size() int {
	return len(f.Content)
}

func (d Directory) Size() int {
	total := 0
	for _, value := range d.files {
		total += value.Size()
	}
	return total
}

func main() {
	file1 := File{Name: "file1.txt", Content: "content1"}
	file2 := File{Name: "file2.jpeg", Content: "content2"}

	subDirectory := Directory{
		Name:  "SubDirectory",
		files: []FileSystemEntry{file1, file2},
	}
	rootDirectory := Directory{
		Name:  "RootDirectory",
		files: []FileSystemEntry{file1, subDirectory},
	}

	fmt.Println(rootDirectory.Size())

}
