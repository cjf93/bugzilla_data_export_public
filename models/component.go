package models

import (
	"strconv"
	"strings"
)

type Component struct {
	ID               int
	Name             string
	product_id       int
	initialowner     int
	initialqacontact int
	description      string
	isactive         bool
}

type ComponentList []Component

type ComponentMap map[string]Component

func CreateComponent(id, product_id, initialowner, initialqacontact int, name, description string, isactive bool) Component {
	var newComponent Component
	newComponent.ID = id
	newComponent.Name = name
	newComponent.product_id = product_id
	newComponent.initialowner = initialowner
	newComponent.initialqacontact = initialqacontact
	str := strings.Replace(description, "\\", "\\\\", -1)
	str = strings.Replace(str, "'", "\\'", -1)
	str = strings.Replace(str, "\"", "\\\"", -1)
	newComponent.description = str
	newComponent.isactive = isactive

	return newComponent
}

func (c *Component) GenerateInsert() string {
	if c.initialqacontact == -1 {
		query := "INSERT INTO `bugs`.`components`(`id`,`name`,`product_id`,`initialowner`,`initialqacontact`,`description`,`isactive`)VALUES(" +
			strconv.Itoa(c.ID) + "," + "\"" + c.Name + "\"" + "," + strconv.Itoa(c.product_id) + "," + strconv.Itoa(c.initialowner) + ", NULL ," +
			"\"" + c.description + "\"" + "," + "\"" + strconv.FormatBool(c.isactive) + "\"" + ");"
		return query
	}
	query := "INSERT INTO `bugs`.`components`(`id`,`name`,`product_id`,`initialowner`,`initialqacontact`,`description`,`isactive`)VALUES(" +
		strconv.Itoa(c.ID) + "," + "\"" + c.Name + "\"" + "," + strconv.Itoa(c.product_id) + "," + strconv.Itoa(c.initialowner) + "," +
		strconv.Itoa(c.initialqacontact) + "," + "\"" + c.description + "\"" + "," + "\"" + strconv.FormatBool(c.isactive) + "\"" + ");"
	return query

	//Succesfull insert
	/*
		INSERT INTO `bugs`.`components`
		(`id`,`name`,`product_id`,`initialowner`,`initialqacontact`,`description`,`isactive`)
		VALUES
		(2,"asd",1,1,NULL,"asd",1);
	*/
}
