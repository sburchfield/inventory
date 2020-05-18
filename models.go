
package main

import (
  "time"

  "github.com/jinzhu/gorm"
)

// stores the version & build info of the app

type envVariables struct {
  appMode                        string
	appPort                        string
	dbSchema                       string
	dbHost                         string
	dbName                         string
	dbUsername                     string
	dbPassword                     string
	localIp                        string
  dbSSLMode                      string
  appPasswordResetDomain         string
  appPasswordResetLinkExpiryTime uint8
}

type user struct {
  gorm.Model
  UserUuid       string
	UserEmail      string
	Username       string
	PasswordHash   string
	FirstName      string
	LastName       string
	Status         string
	PasswordResetHash string
	ResetTime      *time.Time `gorm:"TYPE:timestamp(6) with time zone"`
  Role           string
}

func (user) TableName() string {

	return envVars.dbSchema + ".users"

}

type passResetMessage struct {
  UserUuid    string
	Username    string
	Code        string
	Message     string
}


type Items struct {
  gorm.Model
  ItemName string `gorm:"column:item_name"`
  ItemCost string `gorm:"column:item_cost"`
  ItemPrice string `gorm:"column:item_price"`
  Category string `gorm:"column:category"`
}

func (Items) TableName() string {

	return envVars.dbSchema + ".items"

}

type Stores struct{
  gorm.Model
  StoreName string
  Address    string
  PhoneNumber string
  City       string
  State      string
  ZipCode    string
}

func (Stores) TableName() string {

	return envVars.dbSchema + ".stores"

}

type Categories struct{
  gorm.Model
  Category string
  Description string
}

func (Categories) TableName() string {

	return envVars.dbSchema + ".categories"

}

type Orders struct{
  gorm.Model
  ItemId   string `gorm:"column:item_id"`
  Amount   *int `gorm:"column:amount"`
  UserUuid   string `gorm:"column:user_uuid"`
  StoreId  string   `gorm:"column:store_id"`
}


func (Orders) TableName() string {

	return envVars.dbSchema + ".orders"

}

type EmailOrder struct {
  CreatedAt  *time.Time
  UpdatedAt  *time.Time
  ItemName   string
  Amount     string
  StoreName  string
  FirstName  string
  LastName   string
}
