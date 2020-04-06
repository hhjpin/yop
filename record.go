package yop

import "github.com/jinzhu/gorm"

type Record struct {
	Conn *gorm.DB
}

func (*Record) addRecord(request *Request) error {

	return nil
}

func (*Record) updateRecord(request *Request, response *Response, status string) error {

	return nil
}
