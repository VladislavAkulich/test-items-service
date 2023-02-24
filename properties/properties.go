package properties

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Properties struct {
	App app `yaml:"app"`
	DB  db  `yaml:"db"`
}

type app struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type db struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
}

var once sync.Once
var props Properties

func GetProperties() *Properties {
	once.Do(func() {
		err := cleanenv.ReadConfig("properties.yaml", &props)
		if err != nil {
			panic("Read properties error : " + err.Error())
		}
		fmt.Printf("Properties readed successfully '%v'", props)
	})
	return &props
}
