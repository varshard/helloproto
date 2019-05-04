package addressbook

import (
	"context"
)

type AddressBookServerImpl struct {
	Contacts []*Person
}

func (ads *AddressBookServerImpl) AddPerson(c context.Context, person *Person) (*AddPersonResponse, error) {
	ads.Contacts = append(ads.Contacts, person)

	resp := AddPersonResponse{
		Message: "contact added",
		Success: true,
	}

	return &resp, nil
}

func (ads *AddressBookServerImpl) ListAddressBook(c context.Context, in *ListAddressBookRequest) (*ListAddressBookResponse, error) {
	return &ListAddressBookResponse{
		Success: true,
		Persons: ads.Contacts,
	}, nil
}
