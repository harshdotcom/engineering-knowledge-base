package staff

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

var OverPaidLimit int = 3000

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) OverPaid() []Employee {
	var OverPaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary > OverPaidLimit {
			OverPaid = append(OverPaid, x)
		}
	}

	return OverPaid
}
