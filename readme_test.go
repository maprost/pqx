package pqx_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqtest"
)

func TestMailUsage(t *testing.T) {
	assert := pqtest.InitDatabaseTest(t)
	assert.Nil(MailUsage())
}

func MailUsage() error {
	var err error
	type Mail struct {
		Id         int64     `pqx:"PK AI"`   // is a primary key and use auto increment
		ChangeDate time.Time `pqx:"Changed"` // set automatically by insert and update
		Msg        string
		UserId     int64
	}
	// create a new table 'Mail' to database
	err = pqx.Create(Mail{})
	if err != nil {
		return err
	}

	// insert an entity to database
	mail := Mail{Msg: "Hello world.", UserId: 42}
	err = pqx.Insert(&mail)
	if err != nil {
		return err
	}
	fmt.Println("Id: ", mail.Id)              // '1'
	fmt.Println("Changed: ", mail.ChangeDate) // 'time.now()'

	// select single entity by primary key
	mailCopy := Mail{Id: mail.Id}
	ok, err := pqx.Select(&mailCopy)
	if err != nil {
		return err
	}
	if ok == false {
		return fmt.Errorf("Can't find mail with id: %v", mail.Id)
	}
	fmt.Println("Msg: ", mail.Msg)       // 'Hello world.'
	fmt.Println("UserId: ", mail.UserId) // '42'

	// update an entity by primary key
	mailCopy.Msg = "Bonjour monde."
	err = pqx.Update(&mailCopy)
	if err != nil {
		return err
	}

	// select all mails
	mailList := []Mail{}
	prototype := Mail{}
	pqx.SelectList(&prototype, func() {
		mailList = append(mailList, prototype)
	})
	fmt.Println("Size: ", len(mailList))  // '1'
	fmt.Println("Id: ", mailList[0].Id)   // '1'
	fmt.Println("Msg: ", mailList[0].Msg) // 'Bonjour monde.'

	// delete mail by primary key
	err = pqx.Delete(mail)
	if err != nil {
		return err
	}

	// contains mail by primary key
	contains, err := pqx.Contains(mail)
	if err != nil {
		return err
	}
	fmt.Println("Contains: ", contains) // 'false'
	return nil
}
