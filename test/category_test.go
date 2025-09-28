package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"pakyus_commerce/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCategories(t *testing.T) {
	ClearAll()

	// Categories sudah di-seed melalui migration, jadi tidak perlu seeding manual

	request := httptest.NewRequest(http.MethodGet, "/api/categories", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	var webResponse model.WebResponse[[]model.CategoryResponse]
	err = json.Unmarshal(bytes, &webResponse)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, webResponse.Data)
	assert.Equal(t, 9, len(webResponse.Data)) // Total 9 categories dari migration

	// Check if main categories are returned correctly
	categoryNames := make(map[string]bool)
	parentCategoryNames := make(map[string]bool)
	for _, category := range webResponse.Data {
		categoryNames[category.Name] = true
		if category.ParentId != nil {
			// This is a child category
			parentCategoryNames[category.Name] = true
		}
	}

	// Check main parent categories
	assert.True(t, categoryNames["Elektronik"])
	assert.True(t, categoryNames["Fashion"])
	assert.True(t, categoryNames["Kesehatan"])
	assert.True(t, categoryNames["Makanan & Minuman"])
	assert.True(t, categoryNames["Olahraga"])

	// Check child categories
	assert.True(t, categoryNames["Handphone"])
	assert.True(t, categoryNames["Laptop"])
	assert.True(t, categoryNames["Baju Pria"])
	assert.True(t, categoryNames["Baju Wanita"])

	// Check that child categories have parent_id
	assert.True(t, parentCategoryNames["Handphone"])
	assert.True(t, parentCategoryNames["Laptop"])
	assert.True(t, parentCategoryNames["Baju Pria"])
	assert.True(t, parentCategoryNames["Baju Wanita"])
}
