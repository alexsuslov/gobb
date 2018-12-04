package sqlint

import (
	//_ "database/sql"
	"database/sql/driver"
	"log"
)
type NullInt struct {
	Int int
	Valid bool // Valid is true if Int is not NULL
}

func (n *NullInt) Scan(value interface{}) error {
	if value == nil {
		n.Int, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Int, value)
}


func (n NullInt) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int, nil
}

func convertAssign(dest, src interface{}) error {
	switch s := src.(type) {
	default:
		log.Println("s: ", s)
	}
	return nil
}