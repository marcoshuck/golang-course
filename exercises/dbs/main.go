package main

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name string
}

func main() {
	var db *gorm.DB

	// SELECT * from persons WHERE name LIKE "Marcos" LIMIT 10 ORDER BY id;
	var result []Person

	err := db.Model(&Person{}).Where("name LIKE ?", "Marcos").Limit(10).Order("id DESC").Find(&result).Error
	if err != nil {

	}

}
