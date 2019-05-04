package addressbook

import (
	"context"
	"time"
)

func AddPerson(client AddressBookClient, person Person) (*AddPersonResponse, error) {
	ctx, cancel := makeContext()
	defer cancel()

	// ads.Contacts = append(ads.Contacts, person)

	return client.AddPerson(ctx, &person)
}

func ListAddressBook(client AddressBookClient) (*ListAddressBookResponse, error) {
	ctx, cancel := makeContext()
	defer cancel()

	req := ListAddressBookRequest{}
	return client.ListAddressBook(ctx, &req)
}

func makeContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
