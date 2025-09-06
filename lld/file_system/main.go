package main

import (
	"fmt"
	"path"
	"strings"
)

// === Filter struct ===
type SearchFilter struct {
	Name      string // optional
	Extension string // optional, only for files
}

// === Composite Interface ===
type FileSystemNode interface {
	Name() string
	Display(indent int)
	IsDir() bool
	Search(filter SearchFilter) []FileSystemNode
}

// === File ===
type File struct {
	Filename string
}

func (f *File) Name() string {
	return f.Filename
}

func (f *File) Display(indent int) {
	fmt.Printf("%sFile: %s\n", strings.Repeat(" ", indent), f.Filename)
}

func (f *File) IsDir() bool {
	return false
}

func (f *File) Search(filter SearchFilter) []FileSystemNode {
	var result []FileSystemNode
	if filter.Name != "" && f.Filename != filter.Name {
		return result
	}
	if filter.Extension != "" && !strings.HasSuffix(f.Filename, filter.Extension) {
		return result
	}
	return append(result, f)
}

// === Directory ===
type Directory struct {
	Dirname  string
	Children []FileSystemNode
}

func (d *Directory) Name() string {
	return d.Dirname
}

func (d *Directory) IsDir() bool {
	return true
}

func (d *Directory) Display(indent int) {
	fmt.Printf("%sDirectory: %s\n", strings.Repeat(" ", indent), d.Dirname)
	for _, child := range d.Children {
		child.Display(indent + 2)
	}
}

func (d *Directory) Add(child FileSystemNode) {
	d.Children = append(d.Children, child)
}

func (d *Directory) Search(filter SearchFilter) []FileSystemNode {
	var results []FileSystemNode

	// Verifica se o próprio diretório bate com o filtro de nome
	if filter.Name != "" && d.Dirname == filter.Name {
		results = append(results, d)
	}

	for _, child := range d.Children {
		results = append(results, child.Search(filter)...)
	}
	return results
}

type FileSystemManager struct {
	Root *Directory
}

var instance *FileSystemManager

func GetFileSystemManager() *FileSystemManager {
	if instance == nil {
		instance = &FileSystemManager{
			Root: &Directory{Dirname: "root"},
		}
	}
	return instance
}

// === Public search by path ===
func (fsm *FileSystemManager) Search(pathStr string, filter SearchFilter) []FileSystemNode {
	start := fsm.findPath(pathStr)
	if start == nil {
		fmt.Println("Invalid path:", pathStr)
		return nil
	}
	return start.Search(filter)
}

// === Internal path resolution ===
func (fsm *FileSystemManager) findPath(pathStr string) FileSystemNode {
	cleanedPath := path.Clean(pathStr)
	if cleanedPath == "/" || cleanedPath == "" || cleanedPath == "/root" {
		return fsm.Root
	}

	segments := strings.Split(cleanedPath, "/")[1:] // remove empty string before first '/'

	// Ignora o primeiro segmento "root" pois fsm.Root já representa isso
	if len(segments) > 0 && segments[0] == "root" {
		segments = segments[1:]
	}

	curr := FileSystemNode(fsm.Root)

	for _, seg := range segments {
		if !curr.IsDir() {
			return nil
		}
		found := false
		dir := curr.(*Directory)
		for _, child := range dir.Children {
			if child.Name() == seg {
				curr = child
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}
	return curr
}

func main() {
	fs := GetFileSystemManager()
	root := fs.Root

	docs := &Directory{Dirname: "Documents"}
	pics := &Directory{Dirname: "Pictures"}

	file1 := &File{Filename: "Resume.pdf"}
	file2 := &File{Filename: "Notes.txt"}
	file3 := &File{Filename: "photo.jpg"}

	docs.Add(file1)
	docs.Add(file2)
	pics.Add(file3)

	root.Add(docs)
	root.Add(pics)

	root.Display(0)

	fmt.Println("\n--- Search for .pdf files under /root/Documents ---")
	results := fs.Search("/root/Documents", SearchFilter{Extension: ".pdf"})
	for _, r := range results {
		fmt.Println("Found:", r.Name())
	}

	fmt.Println("\n--- Search for Pictures directory ---")
	dirResults := fs.Search("/root", SearchFilter{Name: "Pictures"})
	for _, r := range dirResults {
		fmt.Println("Found dir:", r.Name())
	}
}
