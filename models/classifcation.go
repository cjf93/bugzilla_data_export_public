package models

import "strconv"

type Classification struct {
	ID   int
	Name string
}

type ClassificationList []Classification

type ClassificationMap map[string]Classification

func CreateClassification(id int, name string) Classification {
	var newClassification Classification
	newClassification.ID = id
	newClassification.Name = name

	return newClassification
}

func (c *Classification) GenerateInsert() string {
	query := "INSERT INTO bugs.classifications	(`id`,`name`) VALUES (" + strconv.Itoa(c.ID) + "," + "\"" + c.Name + "\"" + ");"

	return query

	//Succesfull insert
	/*
		INSERT INTO `bugs`.`classifications`
		(`id`,`name`)
		VALUES
		(2,"test");
	*/
}
