package Control

import (
	"convertPlus/main/Context"
	"convertPlus/main/tag"
)

type Controller struct {
	Context Context.Context
}

func (c Controller) Control(source any, target any, chain []tag.ConvertTag) {

	for _, convertTag := range chain {
		field := convertTag.FieldMetaData.SourceField
	}
}
