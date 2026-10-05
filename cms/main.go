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
func addContact(name string, email string, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Println("\nContact already exists....")
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
	contactIndexByName[name] = len(contactList) - 1
	fmt.Printf("\nContact Added : %s", name)
}
func findContact(name string) *Contact {
	if index, found := contactIndexByName[name]; found {
		return &contactList[index]
	}
	return nil

}
func ListContacts() {
	fmt.Println("\n----CONTACTS-----")
	if len(contactList) > 0 {
		for index, contactItem := range contactList {
			fmt.Printf("\n[%d]. %s", index, contactItem.Name)
		}
	} else {
		fmt.Printf("\nOh no....")
	}
}
func main() {

	ListContacts()
	addContact("nameed", "@eded", "1234")
	addContact("nirmal", "@ed", "34")
	ListContacts()
	addContact("nirmal", "@ed", "34")
	mycontact := findContact("nirmal")
	if mycontact != nil {
		fmt.Printf("\n Found : %+v", *mycontact)
	} else {
		fmt.Println("not foound")
	}

}
