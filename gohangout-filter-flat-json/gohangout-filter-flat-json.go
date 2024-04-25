package main

import (
	"github.com/childe/gohangout/topology"
)

func (f *SplitFilter) SetBelongTo(next topology.Processor) {
	f.next = next
}

type SplitFilter struct {
	config            map[interface{}]interface{}
	field             string
	dropOriginalEvent bool
	deepCopy          bool
	next              topology.Processor
}

func New(config map[interface{}]interface{}) interface{} {
	p := &SplitFilter{
		config: config,
	}

	if v, ok := config["field"]; ok {
		p.field = v.(string)
	} else {
		panic("field must be set in Split filter plugin")
	}

	if dropOriginalEvent, ok := config["drop_original_event"]; ok {
		p.dropOriginalEvent = dropOriginalEvent.(bool)
	} else {
		p.dropOriginalEvent = true
	}

	return p
}

func (p *SplitFilter) copyEvent(event map[string]interface{}) map[string]interface{} {
	newEvent := make(map[string]interface{})
	for k, v := range event {
		newEvent[k] = v
	}
	return newEvent
}

func (p *SplitFilter) Filter(event map[string]interface{}) (map[string]interface{}, bool) {
	v := event[p.field]
	if v == nil {
		return event, false
	}

	s, ok := v.([]interface{})
	if !ok {
		return event, false
	}
	for _, part := range s {
		newEvent := p.copyEvent(event)
		for k, v := range part.(map[string]interface{}) {
			newEvent[k] = v
		}
		p.next.Process(newEvent)
	}

	if p.dropOriginalEvent {
		return nil, false
	}
	return event, true
}
