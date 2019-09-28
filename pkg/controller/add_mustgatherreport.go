package controller

import (
	"github.com/masayag/must-gather-operator/pkg/controller/mustgatherreport"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, mustgatherreport.Add)
}
