package factory

import "fmt"

const Hi_FMT string = "Hi, %s"
const HELLO_FMT string = "Hello, %s"

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct {
}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf(Hi_FMT, name)
}

type helloAPI struct {
}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf(HELLO_FMT, name)
}
