package expr

import (
	"reflect"
)

// Context tracks arguments added to it and can be used to find arguments or fields of arguments by addresses.
type context struct {
	args []any
}

func (c *context) add(arg any) {
	c.args = append(c.args, arg)
}

func (c *context) find(target any) (argIndex int, argType reflect.Type, fieldPath []string) {
	targetAddr := reflect.ValueOf(target).Pointer()
	targetType := reflect.ValueOf(target).Elem().Type()

	for i, arg := range c.args {
		rv := reflect.ValueOf(arg)
		rt := rv.Elem().Type()

		argStartAddr := rv.Pointer()
		argEndAddr := argStartAddr + rt.Size()

		if targetAddr == argStartAddr && rt == targetType {
			return i, rt, nil
		}

		if targetAddr >= argStartAddr && targetAddr < argEndAddr {
			argIndex = i
			argType = rt
			parent := rv.Elem()

		outer:
			for {
				for j := 0; j < parent.NumField(); j++ {
					field := parent.Field(j)
					fieldStartAddr := field.Addr().Pointer()
					fieldEndAddr := fieldStartAddr + field.Type().Size()

					if targetAddr == fieldStartAddr && field.Type() == targetType {
						fieldPath = append(fieldPath, parent.Type().Field(j).Name)
						return
					} else if targetAddr >= fieldStartAddr && targetAddr < fieldEndAddr {
						fieldPath = append(fieldPath, parent.Type().Field(j).Name)
						parent = field
						continue outer
					}
				}
				return -1, nil, nil
			}
		}
	}

	return -1, nil, nil
}

func find(target any) (argIndex int, argType reflect.Type, fieldPath []string) {
	for ctx := range contexts.data {
		argIndex, argType, fieldPath = ctx.find(target)
		if argIndex >= 0 {
			return
		}
	}
	return -1, nil, nil
}
