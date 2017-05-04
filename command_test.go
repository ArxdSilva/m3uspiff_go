package main

import (
	"os"
	"testing"
)

func TestCommandOneArg(t *testing.T) {
	os.Args = []string{"one", ""}
	command()
	if err == 1 {
		t.Fail()
	}
}

func TestCommandNoArg(t *testing.T) {
	os.Args = []string{}
	command()
	if err != 1 {
		t.Fail()
	}
}

func TestCommandTwoArg(t *testing.T) {
	os.Args = []string{"one", "two"}
	command()
	if err != 1 {
		t.Fail()
	}
}
