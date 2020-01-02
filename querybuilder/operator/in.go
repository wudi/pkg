package operator

import "reflect"

func init() {
	AddOperator(In)
}

// In operator check if value is within input param
var In = &Operator{
	Name: "in",
	Evaluate: func(input, value interface{}) bool {
		rinput := reflect.ValueOf(input)

		if rinput.Kind() == reflect.Slice {
			for i := 0; i < rinput.Len(); i++ {
				if rinput.Index(i).Interface() == value {
					return true
				}
			}
		}

		return false
	},
}