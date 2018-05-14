package layout

import (
	"encoding/json"
	"fmt"
	"reflect"
	"screw/log/event"
)

type LayoutConf interface {
	TypeName() string
	Marshal_JSON() ([]byte, error)
	Unmarshal_JSON([]byte) error
}

type Layout interface {
	Format(event *event.Event) string
	LoadConf(LayoutConf) error
}

// Factory 记录conf和layout的绑定关系
type Factory struct {
	apps map[string]App
}

type App struct {
	conf   LayoutConf
	worker Layout
}

type BaseConf struct {
	Type string
}

func NewFactory() *Factory {
	factory := &Factory{
		apps: make(map[string]App),
	}

	return factory
}

func (this *Factory) Register(conf LayoutConf, worker Layout) {
	app := App{
		conf:   conf,
		worker: worker,
	}

	this.apps[conf.TypeName()] = app
}

// ParseJSON 解析JSON获取对应的layout
func (this *Factory) ParseJSON(data []byte) (worker Layout, err error) {
	baseconf := &BaseConf{}

	err = json.Unmarshal(data, baseconf)
	if err != nil {
		return
	}

	app, exist := this.apps[baseconf.Type]
	if !exist {
		err = fmt.Errorf("LayoutFactory can not find type[%s].", baseconf.Type)
		return
	}

	confRefVal := reflect.New(reflect.TypeOf(app.conf).Elem())
	conf, ok := confRefVal.Interface().(LayoutConf)
	if !ok {
		err = fmt.Errorf("LayoutFactory convert [%T] to LayoutConf fail.", confRefVal.Elem().Interface())
		return
	}

	workerRefVal := reflect.New(reflect.TypeOf(app.worker).Elem())
	worker, ok = workerRefVal.Interface().(Layout)
	if !ok {
		err = fmt.Errorf("LayoutFactory convert [%T] to Layout fail.", workerRefVal.Elem().Interface())
		return
	}

	// parse data
	err = conf.Unmarshal_JSON(data)
	if err != nil {
		return
	}

	err = worker.LoadConf(conf)
	if err != nil {
		return
	}

	return
}

// global factory
var (
	globalFactory *Factory
)

func init() {
	globalFactory = NewFactory()

	// default register
	globalFactory.Register(&SimpleLayoutConf{}, &SimpleLayout{})
}

func Register(conf LayoutConf, worker Layout) {
	globalFactory.Register(conf, worker)
}

func ParseJSON(data []byte) (worker Layout, err error) {
	return globalFactory.ParseJSON(data)
}
