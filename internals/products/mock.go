package products

import (
	"context"
)

// Mock แทน DB ได้เลย
type MockRepository struct{}

func (m *MockRepository) GetProducts(ctx context.Context) ([]ProductsModel, error) {
    return []ProductsModel{{Idnumber: 1, Name: "mock"},
	{Idnumber: 2, Name: "test"},
	}, nil  // ✅ ไม่ต้องการ DB
}

// func TestGetProducts(t *testing.T) {
//     svc := NewService(&MockRepository{})  // ง่ายมาก
// }