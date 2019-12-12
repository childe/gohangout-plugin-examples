package main

import (
	"strings"

	"github.com/childe/gohangout/field_setter"
	"github.com/childe/gohangout/value_render"
)

// TitleFilter implents github.com/childe/gohangout/filter.Filter interface
type TitleFilter struct {
	config map[interface{}]interface{}
	fields map[field_setter.FieldSetter]value_render.ValueRender
}

// New returns a filter.Filter interface
func New(config map[interface{}]interface{}) interface{} {
	plugin := &TitleFilter{
		config: config,
		fields: make(map[field_setter.FieldSetter]value_render.ValueRender),
	}

	if fieldsValue, ok := config["fields"]; ok {
		for _, f := range fieldsValue.([]interface{}) {
			fieldSetter := field_setter.NewFieldSetter(f.(string))
			if fieldSetter == nil {
				panic("could build field setter from " + f.(string))
			}
			plugin.fields[fieldSetter] = value_render.GetValueRender2(f.(string))
		}
	} else {
		panic("fields must be set in title filter plugin")
	}
	return plugin
}

// Filter titles fields in the config
func (plugin *TitleFilter) Filter(event map[string]interface{}) (map[string]interface{}, bool) {
	for fs, v := range plugin.fields {
		valueI := v.Render(event)
		if value, ok := valueI.(string); ok {
			event = fs.SetField(event, strings.Title(value), "", true)
		}
	}
	return event, true
}
