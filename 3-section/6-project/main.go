package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}
func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already Exist %+v\n", name)
		return
	}
	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1 // newContact.ID
	fmt.Printf("Contact added: %v\n", name)

}

func findContactByName(name string) *Contact {
	index, exists := contactIndexByName[name]
	if exists {
		return &contactList[index]
	}
	return nil
}

func ListContacts() {
	fmt.Println("----- Listing_ Contacts ------")
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	for i, contact := range contactList {
		fmt.Printf("%d. ID: %d, Name: %s, Email: %s, Phone: %s\n", i+1, contact.ID, contact.Name, contact.Email, contact.Phone)
	}
	fmt.Println("")
}

func main() {
	ListContacts()
	addContact("Alice WonderLand", "alice@gmail.com", "900900900")
	addContact("Bob The Builder", "bob@gmail.com", "900900900")
	// addContact("Mark The CEO", "mark@gmail.com", "900900900")
	// addContact("John The Killer", "john@gmail.com", "900900900")
	addContact("James The Agent", "james@gmail.com", "900900900")
	addContact("James The Agent", "james@gmail.com", "900900900") // attempt to add duplicate
	ListContacts()
	// bob2 := findContactByName("Bob2 The Builder")
	// fmt.Println(bob2.Name) panic without nil check

	bob := findContactByName("Bob The Builder")
	if bob == nil {
		fmt.Println("No Bob contact found")
	} else {
		fmt.Println("bob contact found.", bob.Name)
	}

}
