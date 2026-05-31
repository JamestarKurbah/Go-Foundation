package main

import "fmt"

// Composition --> Has-A relationship
// Inheritance -> Is-A relationship
// Car -> is composed of many parts (Engine, Doors)

type Address struct {
	City    string
	Street  string
	State   string
	ZipCode string
}

type Customer struct {
	Name            string
	Email           string
	BillingAddress  Address
	ShippingAddress Address
}

func (a Address) FullAddress() string {
	if a.City == "" && a.Street == "" {
		return "No address provided"
	}
	return fmt.Sprintf("%s, %s, %s, %s", a.City, a.State, a.Street, a.ZipCode)
}

func (c Customer) PrintDetails() {
	fmt.Printf("Customer Name: %s\nEmail: %s\n", c.Name, c.Email)
	fmt.Println("Billing Address: ", c.BillingAddress.FullAddress())
	fmt.Println("Shipping Address: ", c.ShippingAddress.FullAddress())
}

func main() {

	fmt.Println("--------------- Composition ------------------")

	cust1 := Customer{
		Name:  "Gadget Corp",
		Email: "sales@gadgetcorp.com",
		BillingAddress: Address{
			Street:  "123 Tech Road",
			City:    "Innovateville",
			State:   "CA",
			ZipCode: "9036",
		},
		ShippingAddress: Address{
			Street:  "465 Factory Lane",
			City:    "Manufacturicity",
			State:   "NV",
			ZipCode: "9665",
		},
	}
	cust1.PrintDetails()

	fmt.Println("--------Same Address-------")
	mainAddress := Address{
		Street:  "98 Street",
		City:    "Ahmedabad",
		State:   "Gujarat",
		ZipCode: "75001",
	}
	cust2 := Customer{
		Name:            "Travel Bug",
		Email:           "travelbugcorp@travel.com",
		ShippingAddress: mainAddress,
		BillingAddress:  mainAddress,
	}
	cust2.PrintDetails()
}
