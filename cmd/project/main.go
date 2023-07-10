package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wisnuragaprawida/project/bootstrap"
	cmd "github.com/wisnuragaprawida/project/cmd/project/commands"
	"github.com/wisnuragaprawida/project/pkg/log"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	var (
		err error
	)

	//service dependency
	dependency := bootstrap.NewDependency()
	if err = cmd.Run(dependency); err != nil {
		log.Errorf("unable to execute root command: %s", err)
	}

}
