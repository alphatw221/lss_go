package db

import (
	scylladb "lss_go/lib/db/scylladb"
)



func Connect(){
	scylladb.Connect()
}
func Close(){
	scylladb.Close()
}