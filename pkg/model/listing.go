package model

import "time"

type Listing struct {
	id int32 ``
	name string ``
	desc string
	startprice int
	leadprice int
	starttime time.Time 
	endtime	time.Time 
	createdtime time.Time
	updatedtime time.Time
}
