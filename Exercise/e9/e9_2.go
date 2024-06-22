package main

type MyErr struct {
	Employee string
}

func (me MyErr) Error() string {
	return me.Employee
}
