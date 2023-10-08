package api

import (
	"log"

	"github.com/go-xorm/xorm"
	"github.com/nasa/planetary_tourism/driver"
)

type Env struct {
	orm *xorm.Engine
}

var env = &Env{}

func GetEnv() *Env {
	return env
}

// initialize env when launched server

func InitOrm() *xorm.Engine {
	var err error
	env.orm, err = driver.NewOrm()
	if err != nil {
		log.Fatalf("Init orm error : %s", err)
	}

	return env.orm

}
