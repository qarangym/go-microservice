package models

import "testing"

func TestItem(t *testing.T) {
	item := Item{ID: "1", Name: "TestItem", Price: 100}

	if item.ID != "1" || item.Name != "TestItem" || item.Price != 100 {
		t.Errorf("item creation failed")
	}
}
