package main

type DotInput struct{}

func New(config map[interface{}]interface{}) interface{} {
	return &DotInput{}
}

func (p *DotInput) ReadOneEvent() map[string]interface{} {
	return map[string]interface{}{"message": "."}
}

func (p *DotInput) Shutdown() {}
