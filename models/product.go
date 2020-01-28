package models

import (
	"strconv"
	"strings"
)

type Product struct {
	id                 int
	name               string
	classification_id  int //DEFAULT '1',
	description        string
	isactive           bool   //DEFAULT 1
	defaultmilestone   string //DEFAULT '---',
	allows_unconfirmed bool   //DEFAULT '1',
}

type ProductList []Profile

type ProductMap map[string]Profile

func CreateProduct(id, classification_id int, name, description, defaultmilestone string, isactive, allows_unconfirmed bool) Product {
	var newProduct Product
	newProduct.id = id
	newProduct.classification_id = classification_id
	newProduct.name = name
	str := strings.Replace(description, "\\", "\\\\", -1)
	str = strings.Replace(str, "'", "\\'", -1)
	str = strings.Replace(str, "\"", "\\\"", -1)
	newProduct.description = str
	newProduct.isactive = isactive
	newProduct.defaultmilestone = defaultmilestone
	newProduct.allows_unconfirmed = allows_unconfirmed

	return newProduct
}

func (p *Product) GenerateInsert() string {
	query := "INSERT INTO bugs.products	(`id`,`name`,`classification_id`,`description`,`isactive`,`defaultmilestone`,`allows_unconfirmed`) VALUES (" +
		strconv.Itoa(p.id) + "," + "\"" + p.name + "\"" + "," + strconv.Itoa(p.classification_id) + "," +
		"\"" + p.description + "\"" + "," + "\"" + strconv.FormatBool(p.isactive) + "\"" + "," + "\"" + p.defaultmilestone + "\"" + "," +
		"\"" + strconv.FormatBool(p.allows_unconfirmed) + "\"" + ");"

	return query
}
