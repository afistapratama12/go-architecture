package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/afistapratama12/go-architecture/maker"
)

var (
	command = `
list of command go-architecture

init project:
install init <name>        -- install init "go mod init" project
install package project    -- installing REST API package GIN-GONIC and GORL
install driver <name>      -- installing driver database for gorm package eg:mysql / postgresql
install config <name>      -- create configuratiion file with driver name for gorm

create :
create entity <name>       -- for creating all files with entity name
create repo-crud <name>    -- for create syntax CRUD repository with entity name
create service-crud <name> -- for create syntax CRUD service with entity name
create auth                -- for create syntax auth service
create middleware          -- for create syntax middleware for handler

other command:
help 	                   -- show list of go-architecture command
exit                       -- exit go-architecture command`
)

func Command() {

	var path string
	var initNameProject string

	fmt.Println("=================================================")
	fmt.Println("======== WELCOME TO GOLANG ARCHITECTURE =========")
	fmt.Println("=================================================")

	fmt.Print("please insert a specific path: ")
	fmt.Scanf("%s", &path)
	fmt.Print("please insert a specific project: ")
	fmt.Scanf("%s", &initNameProject)

	for {
		split := func() []string {
			fmt.Println("")
			fmt.Print("go-architecture: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			nameCommand := scanner.Text()

			return strings.Split(nameCommand, " ")
		}()

		if split[0] == "install" {
			dict := maker.NewDict()
			createDir := maker.NewCreateDir(path, split[1])
			createFile := maker.NewCreateFile(path, split[1], dict)

			if split[1] == "package" && split[2] == "project" {
				CMD(path, "go get -u github.com/gin-gonic/gin")
				CMD(path, "go get -u gorm.io/gorm")
				CMD(path, "go get -u github.com/dgrijalva/jwt-go")
				CMD(path, "go get -u github.com/gin-contrib/cors")
				continue
			}

			if split[1] == "init" && len(split[2]) > 0 {
				init := fmt.Sprintf("go mod init %s", split[2])
				nameInit := &initNameProject
				*nameInit = split[2]
				CMD(path, init)
				createFile.CreateMain()
				continue

			} else if split[1] == "driver" {
				if split[2] == "mysql" {
					CMD(path, "go get -u gorm.io/driver/mysql")
				} else if split[2] == "postgresql" {
					CMD(path, "go get -u gorm.io/driver/postgres")
				} else if split[2] == "mysql" {
					CMD(path, "go get -u gorm.io/driver/sqlite")
				} else if split[2] == "sqlserver" {
					CMD(path, "go get -u gorm.io/driver/sqlserver")
				} else {
					fmt.Println("please insert a specific driver / dialect database")
				}
				continue

			} else if split[1] == "config" {
				createDir.CreateDirConfig()

				if len(split) == 2 {
					fmt.Println("please insert specific driver for create config file")
				} else {
					createFile.CreateConfig(split[2])
				}
				continue
			} else {
				fmt.Println("please insert a specific command for install eg: install package project")
			}

		} else if split[0] == "create" {
			dict := maker.NewDict()
			createDir := maker.NewCreateDir(path, split[1])
			createFile := maker.NewCreateFile(path, split[1], dict)

			if split[1] == "entity" {
				createDir.CreateDirectory()
				createDir.CreateDirHandler()

				createFile.CreateModel()
				createFile.CreateRepository()
				createFile.CreateService()
				createFile.CreateHandler(initNameProject)
				continue
			} else if split[1] == "repo-crud" {
				createDir.CreateDirectory()

				createFile.CreateCRUDRepository()
				continue
			} else if split[1] == "service-crud" {
				createDir.CreateDirectory()

				createFile.CreateCRUDService()
				continue
			} else if split[1] == "auth" {
				createDir.CreateDirAuth()

				createFile.CreateAuthService()
				continue
			} else if split[1] == "middleware" {
				createDir.CreateDirMiddleware()

				createFile.CreateMiddleware(initNameProject)
				continue
			} else {
				fmt.Println("please insert a specific 'create' command ")
			}

		} else if split[0] == "help" {
			fmt.Println(command)
			continue
		} else if split[0] == "exit" {
			break
		} else {
			fmt.Println("error command : please insert correct command")
		}
	}
}
