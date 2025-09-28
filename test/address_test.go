package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"pakyus_commerce/internal/entity"
	"pakyus_commerce/internal/model"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAddress(t *testing.T) {
	// Create test user
	user := CreateTestUserWithLogin(t)

	// Create test contact
	contact := &entity.Contact{
		ID:        uuid.New(),
		FirstName: "Test",
		LastName:  "Contact",
		Email:     "test@example.com",
		Phone:     "08123456789",
		UserId:    user.ID,
	}
	err := db.Create(contact).Error
	assert.Nil(t, err)

	requestBody := model.CreateAddressRequest{
		Street:     "Jalan Belum Jadi",
		City:       "Jakarta",
		Province:   "DKI Jakarta",
		PostalCode: "343443",
		Country:    "Indonesia",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/contacts/"+contact.ID.String()+"/addresses", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.AddressResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Street, responseBody.Data.Street)
	assert.Equal(t, requestBody.City, responseBody.Data.City)
	assert.Equal(t, requestBody.Province, responseBody.Data.Province)
	assert.Equal(t, requestBody.Country, responseBody.Data.Country)
	assert.Equal(t, requestBody.PostalCode, responseBody.Data.PostalCode)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
	assert.NotNil(t, responseBody.Data.ID)
}

func TestCreateAddressFailed(t *testing.T) {
	// Create test user
	user := CreateTestUserWithLogin(t)

	// Create test contact
	contact := &entity.Contact{
		ID:        uuid.New(),
		FirstName: "Test",
		LastName:  "Contact",
		Email:     "test@example.com",
		Phone:     "08123456789",
		UserId:    user.ID,
	}
	err := db.Create(contact).Error
	assert.Nil(t, err)

	requestBody := model.CreateAddressRequest{
		Street:     "Very long street name that exceeds the maximum limit of 255 characters which should trigger a validation error because this street name is intentionally made very long to exceed the validation limit that has been set in the validation rules for the street field in the CreateAddressRequest struct",
		City:       "",
		Province:   "",
		PostalCode: "",
		Country:    "",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/contacts/"+contact.ID.String()+"/addresses", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.AddressResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestListAddresses(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)

	CreateAddresses(t, contact, 5)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts/"+contact.ID.String()+"/addresses", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.AddressResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 5, len(responseBody.Data))
}

func TestListAddressesFailed(t *testing.T) {
	TestCreateContact(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)

	CreateAddresses(t, contact, 5)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts/"+"00000000-0000-0000-0000-000000000000"+"/addresses", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.AddressResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestGetAddress(t *testing.T) {
	TestCreateAddress(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	address := GetFirstAddress(t, contact)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts/"+contact.ID.String()+"/addresses/"+address.ID.String(), nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.AddressResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, address.ID.String(), responseBody.Data.ID)
	assert.Equal(t, address.Street, responseBody.Data.Street)
	assert.Equal(t, address.City, responseBody.Data.City)
	assert.Equal(t, address.Province, responseBody.Data.Province)
	assert.Equal(t, address.Country, responseBody.Data.Country)
	assert.Equal(t, address.PostalCode, responseBody.Data.PostalCode)
	assert.Equal(t, address.CreatedAt, responseBody.Data.CreatedAt)
	assert.Equal(t, address.UpdatedAt, responseBody.Data.UpdatedAt)
}

func TestGetAddressFailed(t *testing.T) {
	TestCreateAddress(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)

	request := httptest.NewRequest(http.MethodGet, "/api/contacts/"+contact.ID.String()+"/addresses/"+"00000000-0000-0000-0000-000000000000", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.AddressResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestUpdateAddress(t *testing.T) {
	TestCreateAddress(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	address := GetFirstAddress(t, contact)

	requestBody := model.CreateAddressRequest{
		Street:     "Jalan Lagi Dijieun",
		City:       "Bandung",
		Province:   "Jawa Barat",
		PostalCode: "343443",
		Country:    "Indonesia",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/contacts/"+contact.ID.String()+"/addresses/"+address.ID.String(), strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.AddressResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Street, responseBody.Data.Street)
	assert.Equal(t, requestBody.City, responseBody.Data.City)
	assert.Equal(t, requestBody.Province, responseBody.Data.Province)
	assert.Equal(t, requestBody.Country, responseBody.Data.Country)
	assert.Equal(t, requestBody.PostalCode, responseBody.Data.PostalCode)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
	assert.NotNil(t, responseBody.Data.ID)
}

func TestUpdateAddressFailed(t *testing.T) {
	TestCreateAddress(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	address := GetFirstAddress(t, contact)

	requestBody := model.UpdateAddressRequest{
		Street:     "",
		City:       "",
		Province:   "",
		PostalCode: "12345678901",
		Country:    "",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/contacts/"+contact.ID.String()+"/addresses/"+address.ID.String(), strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.AddressResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestDeleteAddress(t *testing.T) {
	TestCreateAddress(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)
	address := GetFirstAddress(t, contact)

	request := httptest.NewRequest(http.MethodDelete, "/api/contacts/"+contact.ID.String()+"/addresses/"+address.ID.String(), nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[bool])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, true, responseBody.Data)
}

func TestDeleteAddressFailed(t *testing.T) {
	TestCreateAddress(t)

	user := GetFirstUser(t)
	contact := GetFirstContact(t, user)

	request := httptest.NewRequest(http.MethodDelete, "/api/contacts/"+contact.ID.String()+"/addresses/"+"00000000-0000-0000-0000-000000000000", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[bool])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}
