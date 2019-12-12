package main

type EmptyDecoder struct {
}

func New() interface{} {
	return &EmptyDecoder{}
}

// Decode convert any byte array to a empty map
func (d *EmptyDecoder) Decode([]byte) map[string]interface{} {
	return make(map[string]interface{})
}
