package entity

import "fmt"

type ErrorPropertyNotFound[T PropertyTypes] struct {
	PropertyId PropertyId[T]
}

// Implement the Error() method for the MyError type
func (e *ErrorPropertyNotFound[T]) Error() string {
	return fmt.Sprintf("error property not found in entity propertyId:%d", e.PropertyId)
}

func NewErrorPropertyNorFound[T PropertyTypes](propertyId PropertyId[T]) *ErrorPropertyNotFound[T] {
	return &ErrorPropertyNotFound[T]{propertyId}
}
