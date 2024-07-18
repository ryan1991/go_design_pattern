package visitor

import "testing"

func TestServiceRequestVisitor(t *testing.T) {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("CompanyA"))
	c.Add(NewEnterpriseCustomer("CompanyB"))
	c.Add(NewIndividualCustomer("bob"))
	c.Accept(&ServiceRequestVisitor{})

}

func TestAnalysisVisitor(t *testing.T) {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("CompanyA"))
	c.Add(NewEnterpriseCustomer("CompanyB"))
	c.Add(NewIndividualCustomer("bob"))
	c.Accept(&AnalysisVisitor{})

}
