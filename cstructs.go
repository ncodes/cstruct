package cstructs

import "github.com/fatih/structs"
import "fmt"
import "strings"

var strictMode = true

// SetStrictMode enables case sensitive field comparison while false
// disables case sensitivity.
func SetStrictMode(s bool) {
	strictMode = s
}

// Copy copies the value of fields from src to similar
// fields of dest
func Copy(src interface{}, dest interface{}) error {

	if !structs.IsStruct(src) {
		return fmt.Errorf("src is not a struct")
	} else if !structs.IsStruct(dest) {
		return fmt.Errorf("dest is not a struct")
	}

	_src := structs.New(src)
	_dest := structs.New(dest)
	_srcFields := _src.Fields()
	_destFields := _dest.Fields()

	for _, srcField := range _srcFields {
		for _, destField := range _destFields {
			if strictMode {
				if srcField.Name() == destField.Name() && srcField.Kind().String() == destField.Kind().String() {
					destField.Set(srcField.Value())
				}
			} else {
				if strings.ToLower(srcField.Name()) == strings.ToLower(destField.Name()) && strings.ToLower(srcField.Kind().String()) == strings.ToLower(destField.Kind().String()) {
					destField.Set(srcField.Value())
				}
			}
		}
	}

	return nil
}
