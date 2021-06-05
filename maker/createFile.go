package maker

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type CreateFile interface {
	CreateFile(pathDir string, data string) error
	CreateMain()
	CreateRepository()
	CreateService()
	CreateModel()

	CreateConfig(driver string)
	CreateHandler(nameMod string)

	// ini belum dibuat
	CreateCRUDRepository()
	CreateCRUDService()

	CreateAuthService()
	CreateMiddleware()
}

type createFile struct {
	path string
	dir  string
	dict Dict
}

func NewCreateFile(path string, dir string, dict Dict) *createFile {
	return &createFile{path, dir, dict}
}

func (c *createFile) CreateFile(pathDir string, data string) error {
	_, err := os.Stat(pathDir)
	if os.IsExist(err) {
		return err
	}

	f, err := os.Create(pathDir)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err2 := f.WriteString(data)

	if err2 != nil {
		return err2
	}

	return nil
}

func (c *createFile) CreateMain() {
	data := c.dict.DictMain()

	err := c.CreateFile(c.path+"/main.go", data)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("success create main.go")
}

// create formater for clean code architecture
// c.dir for name file entity
func (c *createFile) CreateRepository() {
	var data = c.dict.DictRepository(c.dir)

	err := c.CreateFile(c.path+"/"+c.dir+"/repository.go", data)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("success create repository", c.dir)
}

// create formater for clean code architecture
// c.dir for name file entity
func (c *createFile) CreateService() {
	var data = c.dict.DictService(c.dir)

	err := c.CreateFile(c.path+"/"+c.dir+"/service.go", data)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("success create service", c.dir)

}

// create formater for clean code architecture
// c.dir for name file entity
func (c *createFile) CreateModel() {

	var namePascalFirst = strings.ToUpper(string(c.dir[0]))
	var namePascal = namePascalFirst + string(c.dir[1:])

	var data = c.dict.DictModel(c.dir, namePascal)

	err := c.CreateFile(c.path+"/"+c.dir+"/entity.go", data)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("success create entity", c.dir)
}

func (c *createFile) CreateHandler(nameMod string) {
	var namePascalFirst = strings.ToUpper(string(c.dir[0]))
	var namePascal = namePascalFirst + string(c.dir[1:])

	dir := fmt.Sprintf("%s/handler/%sHandler.go", c.path, c.dir)

	data := c.dict.DictHandler(c.dir, namePascal, nameMod)

	err := c.CreateFile(dir, data)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("success create handler entity", c.dir)
}

func (c *createFile) CreateConfig(driver string) {
	if driver == "mysql" {
		data := c.dict.DictConfigMySQL()

		err := c.CreateFile(c.path+"/config/config.go", data)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("success create config for driver mysql")
		return
	}

	if driver == "postgresql" {
		data := c.dict.DictConfigPostgreSQL()

		err := c.CreateFile(c.path+"/config/config.go", data)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("success create config for driver postgresql")
		return
	}

	if driver == "sqlite" {
		data := c.dict.DictConfigSQLite()

		err := c.CreateFile(c.path+"/config/config.go", data)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("success create config for driver sqlite")
		return
	}

	if driver == "sqlserver" {
		data := c.dict.DictConfigSQLServer()

		err := c.CreateFile(c.path+"/config/config.go", data)

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("success create config for driver sqlserver")
		return
	}

	fmt.Println("please insert a correct driver")
}

func (c *createFile) CreateCRUDRepository() {
	var namePascalFirst = strings.ToUpper(string(c.dir[0]))
	var namePascal = namePascalFirst + string(c.dir[1:])

	var data = c.dict.DictCRUDRepository(c.dir, namePascal)

	err := c.CreateFile(c.path+"/"+c.dir+"/repository.go", data)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("success create CRUD repository", c.dir)

}
func (c *createFile) CreateCRUDService() {
	var namePascalFirst = strings.ToUpper(string(c.dir[0]))
	var namePascal = namePascalFirst + string(c.dir[1:])

	var data = c.dict.DictCRUDService(c.dir, namePascal)

	err := c.CreateFile(c.path+"/"+c.dir+"/service.go", data)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("success create CRUD service", c.dir)

}

func (c *createFile) CreateAuthService() {
	// var namePascalFirst = strings.ToUpper(string(c.dir[0]))
	// var namePascal = namePascalFirst + string(c.dir[1:])

	var data = c.dict.DictAuthService()

	err := c.CreateFile(c.path+"/auth/service.go", data)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("success create auth service", c.dir)

}
func (c *createFile) CreateMiddleware() {
	var data = c.dict.DictMiddleware()

	err := c.CreateFile(c.path+"/middleware/middleware.go", data)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("success create middleware", c.dir)

}
