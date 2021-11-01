package models

type AuthorizeInput struct {
	Login    string `json:"login" binding:"required"`
	IP       string `json:"ip" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SubnetInput struct {
	Subnet string `json:"subnet" binding:"required"`
}
