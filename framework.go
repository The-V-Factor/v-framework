package vframework

type Handler interface {
	GetName() string
	SetName() string
}

type NetworkConf struct {
	Tls  bool
	Port int
}

type Framework struct {
	controllers []Handler
	svcs        []Handler
	entities    []Handler

	Http *NetworkConf
	Grpc *NetworkConf
}

func new() Framework {
	return Framework{}
}

func (f Framework) WithControllers(controllers ...Handler) Framework {
	controllerMap := getHandlerMap(f.controllers)
	for _, v := range controllers {
		if _, ok := controllerMap[v.GetName()]; ok {
			//TODO print warn
		} else {
			f.controllers = append(f.controllers, v)
		}
	}
	return f
}

func (f Framework) WithSvcs(svcs ...Handler) Framework {
	svcMap := getHandlerMap(f.svcs)
	for _, v := range svcs {
		if _, ok := svcMap[v.GetName()]; ok {
			//TODO print warn
		} else {
			f.svcs = append(f.svcs, v)
		}
	}
	return f
}

func (f Framework) WithEntities(entities ...Handler) Framework {
	entityMap := getHandlerMap(f.entities)
	for _, v := range entities {
		if _, ok := entityMap[v.GetName()]; ok {
			//TODO print warn
		} else {
			f.entities = append(f.entities, v)
		}
	}
	return f
}

func (f Framework) WithHttp(conf NetworkConf) Framework {
	// TODO check conf
	f.Http = &conf
	return f
}

func (f Framework) WithGrpc(conf NetworkConf) Framework {
	// TODO check conf
	f.Grpc = &conf
	return f
}

func (f Framework) Run() {

}

func getHandlerMap(handles []Handler) map[string]Handler {
	result := map[string]Handler{}
	for _, v := range handles {
		result[v.GetName()] = v
	}
	return result
}
