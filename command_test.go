package main

import (
	"errors"
	"os"
	"testing"
)

func TestCommandOneArg(t *testing.T) {
	os.Args = []string{"one", ""}
	command()
	if err == errors.New("m3uspiff takes exactly ONE argument") {
		t.Fail()
	}
}

func TestCommandNoArg(t *testing.T) {
	os.Args = []string{}
	command()
	if err == nil {
		t.Fail()
	}
}

func TestCommandTwoArg(t *testing.T) {
	os.Args = []string{"one", "two"}
	command()
	if err == nil {
		t.Fail()
	}
}
