package main

import (
	"fmt"
)

// DashOutput implents github.com/childe/gohangout/filter.Filter interface
type DashOutput struct {
	config map[interface{}]interface{}
}

// New returns a output.Output interface
func New(config map[interface{}]interface{}) interface{} {
	return &DashOutput{
		config: config,
	}
}

// Emit output '-' and new line
func (p *DashOutput) Emit(event map[string]interface{}) {
	fmt.Println("-")
}

func (p *DashOutput) Shutdown() {}
