package users

import "logistics/models"

type GetAddressBookResponse struct {
	AddressBook []models.Address `json:"address_book"` // 企业通讯录
}
