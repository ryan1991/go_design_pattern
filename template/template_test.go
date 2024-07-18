package template

import "testing"

func TestTemplate(t *testing.T) {

	sub1 := &Sub1{}
	// sub1.AbstractTemplate = &AbstractTemplate{}
	TemplateMethod(sub1)

	sub2 := &Sub2{}
	TemplateMethod(sub2)

}
