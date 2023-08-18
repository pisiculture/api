package domain

import "errors"

type Installation struct {
	Base
	Name        string
	Description string
}

func NewInstallation() *Installation {
	return &Installation{}
}

func (ins *Installation) Validate() error {

	if ins.Name == "" {
		return errors.New("Param { Name } required")
	}

	if ins.Description == "" {
		return errors.New("Param { Description } required")
	}

	return nil
}
