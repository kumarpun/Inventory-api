package main

type Device struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Project   string `json:"project"`
	Device    string `json:"device"`
	Startdate string `json:"startdate"`
	Enddate   string `json:"enddate"`
	UserId    string `json:"userId"`
	Email     string `json:"email"`
	Status    string `json:"status"`
}
