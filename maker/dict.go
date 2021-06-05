package maker

import "fmt"

type Dict interface {
	DictMain() string

	DictRepository(name string) string
	DictService(name string) string
	DictModel(name string, namePascal string) string
	DictHandler(name string, namePascal string, initMod string) string

	DictConfigMySQL() string
	DictConfigPostgreSQL() string
	DictConfigSQLite() string
	DictConfigSQLServer() string

	// ini belom dibuat
	DictCRUDRepository(name string, namePascal string) string
	DictCRUDService(name string, namePascal string) string

	DictAuthService() string
	DictMiddleware() string
}

type dict struct {
}

func NewDict() *dict {
	return &dict{}
}

func (d *dict) DictMain() string {
	return `package main
	
func main() {
}	
`
}

func (d *dict) DictRepository(name string) string {
	return fmt.Sprintf(`package %s
	
import (
	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
`, name)
}

func (d *dict) DictService(name string) string {
	return fmt.Sprintf(`package %s
	
type Service interface {
}

type service struct {
	repository 	Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
`, name)
}

func (d *dict) DictModel(name string, namePascal string) string {
	return fmt.Sprintf(`package %s
	
import (
	"time"
)

type %s struct {
	%sID      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}	
`, name, namePascal, namePascal)
}

func (d *dict) DictHandler(name string, namePascal string, initMod string) string {
	return fmt.Sprintf(`package handler

// path from name mod/entity	
import(
	"%s/%s"
)

type %sHandler struct {
	service %s.Service
}	

func New%sHandler (service %s.Service) *%sHandler {
	return &%sHandler{service}
}	
`, initMod, name, name, name, namePascal, name, name, name)
}

func (d *dict) DictConfigMySQL() string {
	return `package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	)
	
func Config() *gorm.DB {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}`
}

func (d *dict) DictConfigPostgreSQL() string {
	return `package config
	
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	)
	
func Config() *gorm.DB {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
		
	return db
}`
}

func (d *dict) DictConfigSQLite() string {
	return `package config
	
import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	)
	
func Config() *gorm.DB {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}`
}

func (d *dict) DictConfigSQLServer() string {
	return `package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	)
	
func Config() *gorm.DB {
	// github.com/denisenkom/go-mssqldb
	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}`
}

func (d *dict) DictCRUDRepository(name string, namePascal string) string {
	return fmt.Sprintf(`package %s

import (
	"gorm.io/gorm"
)

type Repository interface{
	FindAll() ([]%s, error)
	FindByID(ID string) (%s, error)
	Create(%s %s) (%s, error)
	Update(ID string, dataUpdate map[string]interface{}) (%s, error) 
	Delete(ID string) (string, error)
}

type repository struct{
	db *gorm.DB
}

func NewService(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]%s, error) {
	var %ss []%s
	
	if err := r.db.Find(&%ss).Error; err != nil {
		return %ss, err
	}

	return %ss, nil
}

func (r *repository) FindByID(ID string) (%s, error) {
	var %s %s

	if err := r.db.Where("id = ?", ID).Find(&%s).Error; err != nil {
		return %s, err
	}

	return %s, nil
}

func (r *repository) Create(%s %s) (%s, error) {
	if err := r.db.Create(&%s).Error; err != nil {
		return %s, err
	}

	return %s, nil
}

func (r *repository) Update(ID string, dataUpdate map[string]interface{}) (%s, error) {

	var %s %s

	if err := r.db.Model(&%s).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return %s, err
	}

	if err := r.db.Where("id = ?", ID).Find(&%s).Error; err != nil {
		return %s, err
	}

	return %s, nil
}

func (r *repository) Delete(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&%s{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

	
`, name, namePascal, namePascal, name, namePascal, namePascal, namePascal, namePascal, name, namePascal, name, name, name, namePascal, name, namePascal, name, name, name, name, namePascal, namePascal, name, name, name, namePascal, name, namePascal, name, name, name, name, name, namePascal)
}

func (d *dict) DictCRUDService(name string, namePascal string) string {
	return fmt.Sprintf(`package %s
	
type Service interface{
	GetAll%s() (%s, error)
	SaveNew%s() (%s, error)
	Get%sByID(ID string) (%s, error)
	Update%sByID(ID string) (%s, error)
	Delete%sByID(ID string) (%s, error)
}

type service struct{
	repository 	Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}


`, name, namePascal, namePascal, namePascal, namePascal, namePascal, namePascal, namePascal, namePascal, namePascal, namePascal)
}

func (d *dict) DictAuthService() string {
	return fmt.Sprintf(`package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var (
	key = ""
)

type Service interface{
	GenerateToken(userID string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtservice struct{
}

func NewService() *jwtservice {
	return &jwtservice{}
}

func (s *jwtService) GenerateToken(ID string) (string, error) {
	claim := jwt.MapClaims{
		"id": ID,
	}

	// generate token using HS256 with claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

`)
}

func (d *dict) DictMiddleware() string {
	return fmt.Sprintf(`package middleware

import (
	"github.com/gin-gonic/gin"
)
	
func Middleware(userService user.Service, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}		
`)
}
