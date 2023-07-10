package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"github.com/wisnuragaprawida/project/pkg/one-go/telemetries/otsql"
)

type Dependency struct {
	// contains filtered or unexported fields
	cfg Config
	db  *sqlx.DB
}

func NewDependency() *Dependency {
	dep := new(Dependency)
	return dep
}

func (dep *Dependency) SetConfig(cfg Config) {
	dep.cfg = cfg
}

func (dep *Dependency) GetConfig() Config {
	return dep.cfg
}

func (dep *Dependency) GetDB() *sqlx.DB {
	if dep.db == nil {
		dep.db = otsql.SqlMustConnect(dep.GetConfig().Database.Write)
	}
	return dep.db
}

func (dep *Dependency) Initialize() error {
	//do something
	return nil
}
