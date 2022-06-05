package mock

import "time"

type Context struct {
}

func (Context) Done() <-chan struct{} {
	return nil
}

func (Context) Err() error {
	return nil
}

func (Context) Deadline() (time.Time, bool) {
	return time.Time{}, true
}

func (Context) Value(interface{}) interface{} {
	return nil
}
