package listener

import "log"

type Handler interface {
	Call(HubMessage)
}

type Logger struct{}

func (l *Logger) Call(msg HubMessage) {
	log.Print(msg)
}

type Registry struct {
	entries []func(HubMessage)
}

func (r *Registry) Add(h func(msg HubMessage)) {
	r.entries = append(r.entries, h)
	return
}

func (r *Registry) Call(msg HubMessage) {
	for _, h := range r.entries {
		go h(msg)
	}
}

func MsgHandlers() Registry {
	var handlers Registry

	handlers.Add((&Logger{}).Call)

	return handlers
}
