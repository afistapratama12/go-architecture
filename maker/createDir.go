package maker

import (
	"fmt"
	"log"
	"os"
)

type CreateDir interface {
	CreateDirectory()
	CreateDirHandler()
	CreateDirConfig()
	CreateDirAuth()
	CreateDirMiddleware()
}

type createDir struct {
	path string
	dir  string
}

func NewCreateDir(path string, dir string) *createDir {
	return &createDir{path, dir}
}

func (d *createDir) CreateDirectory() {
	err := os.Mkdir(d.path+"/"+d.dir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success create directory", d.dir)
}

func (d *createDir) CreateDirHandler() {
	err := os.Mkdir(d.path+"/"+"handler", 0755)
	if err != nil {
		// log.Fatal(err)
		fmt.Println("handler directory already created")
	}

	fmt.Println("success create handler directory")
}

func (d *createDir) CreateDirConfig() {
	err := os.Mkdir(d.path+"/"+"config", 0755)

	if err != nil {
		// log.Fatal(err)
		fmt.Println("config directory already created")
	}

	fmt.Println("success create config directory")
}

func (d *createDir) CreateDirAuth() {
	err := os.Mkdir(d.path+"/"+"auth", 0755)

	if err != nil {
		// log.Fatal(err)
		fmt.Println("auth directory already created")
	}

	fmt.Println("success create auth directory")
}

func (d *createDir) CreateDirMiddleware() {
	err := os.Mkdir(d.path+"/"+"middleware", 0755)

	if err != nil {
		// log.Fatal(err)
		fmt.Println("middleware directory already created")
	}

	fmt.Println("success create middleware directory")
}
