package provider

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxlifetime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

var providers = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session :Register provide is nil")
	}
	if _, dup := providers[name]; dup {
		panic("session:Register called twice for provide " + name)
	}
}
