package cron

type Cron interface {
	AddFuncs() error
}

func NewCrons() []Cron {
	return nil
}

type Engine struct {
	crons []Cron
}

func NewEngine(crons []Cron) *Engine {
	return &Engine{crons: crons}
}

func (e *Engine) Run() error {
	return nil
}
