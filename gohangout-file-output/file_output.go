package main

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/childe/gohangout/codec"
	"github.com/childe/gohangout/value_render"
	"github.com/golang/glog"
)

type fileOutput struct {
	config       map[interface{}]interface{}
	encoder      codec.Encoder
	path         value_render.ValueRender
	overfailFile *os.File

	files map[string]*os.File
}

func New(config map[interface{}]interface{}) interface{} {
	p := &fileOutput{
		config: config,
		files:  make(map[string]*os.File),
	}

	var err error
	if v, ok := config["overfail_path"]; ok {
		p.overfailFile, err = os.OpenFile(v.(string), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			glog.Fatalf("open %s error: %v", v, err)
		}
	} else {
		p.overfailFile, err = ioutil.TempFile("", "")
		if err != nil {
			glog.Fatalf("create temp overfail file error: %v", err)
		}
		glog.Info("create overfail file for file output: %v", p.overfailFile.Name())
	}

	if v, ok := config["codec"]; ok {
		p.encoder = codec.NewEncoder(v.(string))
	} else {
		p.encoder = codec.NewEncoder("json")
	}

	if path, ok := config["path"]; ok {
		p.path = value_render.GetValueRender(path.(string))
	}

	return p

}

func (p *fileOutput) file(event map[string]interface{}) *os.File {
	var ok bool
	var filePath string

	if filePath, ok = p.path.Render(event).(string); !ok {
		return p.overfailFile
	}

	if f, ok := p.files[filePath]; ok {
		return f
	}

	dir := path.Dir(filePath)
	os.MkdirAll(dir, 0755)
	if f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
		glog.Infof("open %s error: %v", filePath, err)
		return p.overfailFile
	} else {
		p.files[filePath] = f
		return f
	}
}

func (p *fileOutput) Emit(event map[string]interface{}) {
	buf, err := p.encoder.Encode(event)
	if err != nil {
		glog.Errorf("marshal %v error: %v", event, err)
	}
	f := p.file(event)
	n, err := f.Write(buf)
	if err != nil {
		glog.Errorf("write file error: %v. original length %d, write length %d", err, len(buf), n)
	}
	f.WriteString("\n")
}

func (p *fileOutput) Shutdown() {
	for _, f := range p.files {
		f.Close()
	}
}
