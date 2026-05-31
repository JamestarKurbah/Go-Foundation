package main

import "fmt"

type Payable interface {
	fmt.Stringer
	CalculatePay() float64
}

type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se SalariedEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12.0
}

func (se SalariedEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual: $%.2f)", se.Name, se.AnnualSalary)
}

type HourlyEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64
}

func (he HourlyEmployee) CalculatePay() float64 {
	return he.HourlyRate * he.HoursWorked
}

func (he HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f/hr, Hours: %.1f)", he.Name, he.HourlyRate, he.HoursWorked)
}

type CommissionEmployee struct {
	Name          string
	BaseSalary    float64
	ComissionRate float64
	SalesAmount   float64
}

func (ce CommissionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.ComissionRate * ce.SalesAmount)
}

func (ce CommissionEmployee) String() string {
	return fmt.Sprintf("Commission: %s (Base: $%.2f, CommRate: %.2f, Sales: $%.2f)", ce.Name, ce.BaseSalary, ce.ComissionRate*100, ce.SalesAmount)
}

func PrintEmployeeSummary[P fmt.Stringer](employee P) {
	fmt.Printf("\n - Processing: %s\n", employee)
}
func ProcessPayroll(employees []Payable) {
	fmt.Println("\n---- Processing Payroll ----")
	totalPayroll := 0.0
	for _, emp := range employees {
		PrintEmployeeSummary(emp)
		pay := emp.CalculatePay()
		fmt.Printf("  Monthly Pay: $%.2f\n", pay)
		totalPayroll += pay
	}
	fmt.Printf("\nTotal Monthly Payroll: $%.2f\n", totalPayroll)
	fmt.Println("-----------------------")
}

func main() {
	fmt.Println("Welcome to the Payroll System!")
	salEmp := SalariedEmployee{Name: "Alice", AnnualSalary: 72000.00}
	hrEmp := HourlyEmployee{Name: "Bob", HourlyRate: 25.00, HoursWorked: 160.0}
	comEmp := CommissionEmployee{Name: "Charlie", BaseSalary: 2000.00, ComissionRate: 0.10, SalesAmount: 15000.00}
	payrolllist := []Payable{
		salEmp, hrEmp, comEmp,
		HourlyEmployee{Name: "Diana", HourlyRate: 30.00, HoursWorked: 150.00},
	}
	ProcessPayroll(payrolllist)
}
