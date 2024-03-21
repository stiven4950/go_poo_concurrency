package main

import "testing"

func TestSum(t *testing.T) {
	/* total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Sum was incorrect, got %d expected %d", total, 11)
	} */

	tables := []struct {
		a int
		b int
		n int
	}{
		{5, 2, 7},
		{5, 8, 13},
		{15, 20, 35},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{5, 2, 5},
		{3, 1, 3},
		{63, 20, 63},
		{80, 99, 99},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)

		if max != item.n {
			t.Errorf("GetMax was incorrect, got %d, expected %d", max, item.n)
		}
	}
}

/* func TestFibonacci(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.a)
		if item.n != fib {
			t.Errorf("Fibonacci was incorrect, got %d, expected %d", fib, item.n)
		}
	}
} */

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (*Employee, error) {
					return &Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (*Person, error) {
					return &Person{
						Name: "Jhon Doe",
						Age:  35,
						DNI:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  35,
					DNI:  "1",
					Name: "Jhon Doe",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	// Almacenar los valores de las funciones originales es importante
	// Para permitir que
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		// Se reemplazan las funciones originales, y esas funciones son las
		// que tomará GetFullTimeEmployeeById
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)
		if err != nil {
			t.Errorf("Error when getting employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
