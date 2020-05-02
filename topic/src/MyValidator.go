package src

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func TopicUrl(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	fmt.Println(topStruct) // 反射
	fmt.Println(topStruct.Interface())
	return false
}