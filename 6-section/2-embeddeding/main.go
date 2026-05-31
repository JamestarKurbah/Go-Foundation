package main

import "fmt"

type Address struct {
	City    string
	Street  string
	State   string
	ZipCode string
}
type ContactInfo struct {
	Email string
	Phone string
}

func (a Address) FullAddress() string {
	if a.City == "" && a.Street == "" {
		return "No address provided"
	}
	return fmt.Sprintf("%s, %s, %s, %s", a.City, a.State, a.Street, a.ZipCode)
}

func (c ContactInfo) DisplayContact() string {
	return fmt.Sprintf("Email: %s, Phone: %s", c.Email, c.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string
}

func (c Company) GetProfile() {
	fmt.Printf("Company Name: %s\n", c.Name)
	fmt.Printf("Location: %s\n", c.FullAddress())
	fmt.Printf("Street (promoted): %s\n", c.Street)
	fmt.Printf("Email: (promoted): %s\n", c.Email)
	fmt.Printf("Business Type: %s\n", c.BusinessType)
}

func main() {
	fmt.Println("--- Struct Embedded ---")
	com1 := Company{
		Name: "Tech Innovators Inc.",
		Address: Address{
			Street:  "123 Innovation Drive",
			City:    "Techville",
			State:   "CA",
			ZipCode: "90210",
		},
		ContactInfo: ContactInfo{
			Email: "info@techinnovators.com",
			Phone: "(555) 123-4567",
		},
		BusinessType: "Technology",
	}
	com1.GetProfile()
	fmt.Printf("\nDirect access to com1.City: %s\n", com1.City)   // Promoted from Address
	fmt.Printf("\nDirect access to com1.Phone: %s\n", com1.Phone) // Promoted from ContactInfo

}
