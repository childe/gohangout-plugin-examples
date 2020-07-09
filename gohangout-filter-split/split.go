package main

import (
	"encoding/json"
	"strings"

	"github.com/childe/gohangout/topology"
)

func (f *SplitFilter) SetBelongTo(next topology.Processor) {
	f.next = next
}

type SplitFilter struct {
	config            map[interface{}]interface{}
	sep               string
	field             string
	dropOriginalEvent bool
	deepCopy          bool
	next              topology.Processor
}

func New(config map[interface{}]interface{}) interface{} {
	p := &SplitFilter{
		config: config,
	}

	if v, ok := config["sep"]; ok {
		p.sep = v.(string)
	} else {
		panic("sep must be set in Split filter plugin")
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

	if deepCopy, ok := config["deep_copy"]; ok {
		p.deepCopy = deepCopy.(bool)
	} else {
		p.deepCopy = true
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

	s, ok := v.(string)
	if !ok {
		return event, false
	}
	parts := strings.Split(s, p.sep)
	if p.deepCopy {
		s, err := json.Marshal(event)
		if err != nil {
			return event, false
		}
		for _, part := range parts {
			newEvent := make(map[string]interface{})
			if json.Unmarshal(s, &newEvent) != nil {
				return event, false
			}
			newEvent[p.field] = part
			p.next.Process(newEvent)
		}
	} else {
		for _, part := range parts {
			newEvent := p.copyEvent(event)
			newEvent[p.field] = part
			p.next.Process(newEvent)
		}
	}

	if p.dropOriginalEvent {
		return nil, false
	}
	return event, true
}
