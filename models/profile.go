package models

import "strconv"

type Profile struct {
	Userid       int
	Login_name   string
	realname     string
	disabledtext string
}

type ProfileList []Profile

type ProfileMap map[string]Profile

func CreateProfile(userid int, login_name, realname string) Profile {
	var newProfile Profile
	newProfile.Userid = userid
	newProfile.Login_name = login_name
	newProfile.realname = realname
	newProfile.disabledtext = ""

	return newProfile
}

func (p *Profile) GenerateInsert() string {
	query := "INSERT INTO bugs.profiles	(`userid`,`login_name`,`realname`,`disabledtext`) VALUES (" + strconv.Itoa(p.Userid) +
		",\"" + p.Login_name + "\",\"" + p.realname + "\",\"" + p.disabledtext + "\");"

	return query
	//Succesfull example
	/*
		INSERT INTO `bugs`.`profiles`
		(`userid`,`login_name`,`realname`,`disabledtext`)
		VALUES
		(2,"randommail@somemail.com","randonname"," ");
	*/
}
