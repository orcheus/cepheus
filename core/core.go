package core

type Context struct {
	Id string
}

type CtxManager struct {
	Contexts      map[string]Context
	Subscriptions []string
}

func NewCtxManager() *CtxManager {
	return &CtxManager{
		Contexts:      make(map[string]Context),
		Subscriptions: make([]string, 0),
	}
}

func (cm *CtxManager) register(ctx Context) error {
	return nil
}
