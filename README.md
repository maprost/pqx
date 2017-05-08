[![Build Status](https://travis-ci.org/maprost/pqx.svg?branch=master)](https://travis-ci.org/maprost/pqx)
[![Coverage Status](https://coveralls.io/repos/github/maprost/pqutil/badge.svg)](https://coveralls.io/github/maprost/pqutil)
[![GoDoc](https://godoc.org/github.com/mleuth/pqlib?status.svg)](https://godoc.org/github.com/mleuth/pqlib)
[![Go Report Card](https://goreportcard.com/badge/github.com/maprost/pqx)](https://goreportcard.com/report/github.com/maprost/pqx)

# pqx v0.2
Small lightweight pq library extension. 

## Features
- log query feature
- simplify default actions (see actions)
- scan a struct 
- sql parameter without $1-$n

## Actions
- insert
- update
- delete
- create
- register
- contains (by id)
- select (by id)
- select (by key-value, select single element)
- select all (select multi elements)
- select all (by key value, select multi elements)

## Supported Types
- Integer:
  - `int8`, `int16`, `int`, `int32`, `int64`
  - `sql.NullInt64` 
  - `pqnull.Int8`,  `pqnull.Int16`, `pqnull.Int32`, `pqnull.Int`, `pqnull.Int64`
- Unsigned Integer:
  - `uint8`, `uint16`, `uint`, `uint32`, `uint64`
- Float:
  - `float32`, `float64`
  - `sql.NullFloat64`
  - `pqnull.Float32`, `pqnull.Float64`
- Boolean:
  - `bool`
  - `sql.NullBool`
  - `pqnull.Bool`
- String:
  - `string`
  - `sql.NullString`
  - `pqnull.String`
- Time:
  - `time.Time`
  - `pq.NullTime`
  - `pqnull.Time`

## Supported Tags
- `pqx:"PK"` -> primary key
- `pqx:"AI"` -> auto increment
    - available for: `int8`, `uint8`, `int16`, `uint16`, `int`, `int32`, `uint`, `uint32`, `int64`
- `pqx:"Created"` -> set initial time automatically
- `pqx:"Changed"` -> set initial and update time automatically

## Usage
```go
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


```