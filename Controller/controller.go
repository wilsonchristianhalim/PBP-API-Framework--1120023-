package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {
	db := Connect()
	defer db.Close()

	result, err := db.Query("SELECT * FROM game_account")
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	var account Account
	var accounts []Account

	for result.Next() {
		err = result.Scan(&account.ID, &account.Username, &account.Password)
		if err != nil {
			panic(err.Error())
		}
		accounts = append(accounts, account)
	}

	if len(accounts) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusCreated, accounts)
	}
}

func AddAccount(c *gin.Context) {
	db := Connect()
	defer db.Close()

	var account Account

	if err := c.Bind(&account); err != nil {
		fmt.Println(err)
		return
	}

	if len(account.Username) == 0 || len(account.Password) == 0 {
		fmt.Println("Masukkan nama & password")
		return
	} else {
		db.Exec("INSERT INTO GAME_ACCOUNT (USERNAME, PASSWORD) VALUES(?,?)", account.Username, account.Password)
		c.IndentedJSON(http.StatusOK, "Success")
	}
}

func DeleteAccount(c *gin.Context) {
	db := Connect()
	defer db.Close()

	id := c.Query("id")

	result, errQuery := db.Exec("DELETE FROM GAME_ACCOUNT WHERE ID=?", id)

	num, _ := result.RowsAffected()

	if errQuery == nil {
		if num == 0 {
			c.AbortWithStatusJSON(400, "Data Not Found!")
		} else {
			c.IndentedJSON(http.StatusOK, "Success")
		}
	}
}

func UpdateAccount(c *gin.Context) {
	db := Connect()
	defer db.Close()

	var account Account

	if err := c.Bind(&account); err != nil {
		fmt.Println(err)
		return
	}
	result, errQuery := db.Exec("UPDATE GAME_ACCOUNT SET USERNAME=?, PASSWORD=? WHERE ID=?", account.Username, account.Password, account.ID)

	num, _ := result.RowsAffected()

	if errQuery == nil {
		if num == 0 {
			c.AbortWithStatusJSON(400, "Data With That ID Not Found!")
		} else {
			c.IndentedJSON(http.StatusOK, "Success")
		}
	}
}
