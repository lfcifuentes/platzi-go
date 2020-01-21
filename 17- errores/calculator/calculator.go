package calculator

import "errors"

// Sum suma dos numeros
func Sum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("El parametro a no es un numero")
	}
	switch b.(type) {
	case string:
		return 0, errors.New("El parametro b no es un numero")
	}
	return (a.(int) + b.(int)), nil
}
