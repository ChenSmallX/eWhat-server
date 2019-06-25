package dbcontroller

import (
	"database/sql"
)

func userBuilder(db *sql.DB, table string) {
	sql := "
	DROP TABLE IF EXIST "+table+";
	CREATE TABLE " + table + "(
		uid INT(10) NOT NULL AUTO_INCREAMENT,
		wxid VARCHAR(20) NOT NULL,
		name VARCHAR(20) NOT NULL,
		PRIMARY KEY (uid) 
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;"



}
