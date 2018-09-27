package main

import (
	"log"

	"upper.io/db.v3/lib/sqlbuilder"
)

func RemoveDuplicateUser(sess sqlbuilder.Database) {
	rows, err := sess.Query("select third_id, wx_id, count(*) as count from e_optometry_user where wx_id is not null group by third_id, wx_id having count > 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var thirdID int64
	var wxID string
	var count int64
	for rows.Next() {
		rows.Scan(&thirdID, &wxID, &count)
		log.Printf("thir_id: %d, wx_id: %s, count: %d", thirdID, wxID, count)
		if count > 1 {
			sess.Exec("delete from e_optometry_user where wx_id = ? and third_id = ? order by id desc limit ? ", wxID, thirdID, count-1)
		}
	}
}
