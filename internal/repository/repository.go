package repository

import "fmt"

var (
	ErrNotFound      = fmt.Errorf("entity not found")
	ErrAlreadyExists = fmt.Errorf("entity already exists")
	ErrRef           = fmt.Errorf("entity miss required reference to another entity")
)
