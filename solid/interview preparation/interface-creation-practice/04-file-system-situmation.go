package interface_creation_practice

type FileSystemEntry interface {
	Size() int
}

type File struct{}
type Directory struct{}

func (f File) Size() int {

}

func (d Directory) Size() int {

}

func main() {

}
