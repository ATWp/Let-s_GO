package tests

import (
	"flag"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var opt = godog.Options{
	ShowStepDefinitions: false,
	Randomize:           0,
	StopOnFailure:       false,
	Strict:              false,
	NoColors:            false,
	Tags:                "",
	Format:              "progress",
	Concurrency:         0,
	Paths:               nil,
	Output:              colors.Colored(os.Stdout),
}

func __init__() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		s.Step('I try to register user with age (\d) and name: (\w)', iTryToRegisterUserWithAgeName)
	}, opt)
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

func iTryToRegisterUserWithAgeName( age int, name string) error {
	return nil
}