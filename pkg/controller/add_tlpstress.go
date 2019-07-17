package controller

import (
	"github.com/jsanda/tlp-stress-operator/pkg/controller/tlpstress"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, tlpstress.Add)
}
