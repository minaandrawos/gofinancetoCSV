package dataconverter

import "fmt"

var (
	ErrEmptyParam       = fmt.Errorf("Tried to extract CSV headers and rows from Empty Parameter")
	ErrIncompatibleType = fmt.Errorf("Value is of incompatible type, supported types are int,float64,bool,int64,Stringer, and GoStringer")
	ErrNilArgument      = fmt.Errorf("Could not convert struct to params, structValue is nil")
	ErrNotStruct        = fmt.Errorf("Could not convert struct to params, structValue is not struct")
)
