package Control

import (
	"convertPlus/main/Context"
	"convertPlus/main/tag"
	"reflect"
)

type Controller struct {
	Context Context.Context
}

const (
	NEXT int = iota
	STOP
	NEXT_ONE
)

func (c Controller) Control(source any, target any, chain []tag.ConvertTag) {

	sourceValue := reflect.ValueOf(source).Elem()
	targetValue := reflect.ValueOf(target).Elem()

	for _, convertTag := range chain {
		sourceFieldValue := sourceValue.FieldByName(convertTag.FieldMetaData.SourceField.Name)

		handler := convertTag.FieldMetaData.FieldHandler

		beforeValue := handler.Before(sourceFieldValue)

		convert := handler.Convert(beforeValue)

		afterValue := handler.After(convert)

		switch targetValue.Type().Kind() {
		case reflect.String:
			targetValue.SetString(afterValue.(string))
		}
		/*
			action := handler.GetAction()

			switch action {
			case NEXT:
				continue
			case NEXT_ONE:
			case STOP:
				return
			}*/

	}
}

type Action struct {
	ActionCode int
}

func (a *Action) GetAction() int {
	return a.ActionCode
}

func (a *Action) Next() {
	a.ActionCode = NEXT
}

func (a *Action) Stop() {
	a.ActionCode = STOP
}

func (a *Action) NextOne() {
	a.ActionCode = NEXT_ONE
}

type ActionControl interface {
	Next()

	Stop()

	NextOne()

	GetAction() int
}
