package template

import "fmt"

type AbstractTemplate struct{}

type IDetail interface {
	DetailPart()
}

func TemplateMethod(detail IDetail) {
	CommonPart()
	detail.DetailPart()
}

func CommonPart() {

	fmt.Println("公共逻辑")
}

// func (at *AbstractTemplate) DetailPart() {

// 	fmt.Println("可变逻辑部分")
// }

type Sub1 struct {
	*AbstractTemplate
}

func (sub *Sub1) DetailPart() {
	fmt.Println("Sub1逻辑")
}

type Sub2 struct {
	*AbstractTemplate
}

func (sub *Sub2) DetailPart() {
	fmt.Println("Sub2逻辑")
}
