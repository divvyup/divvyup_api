package db

/*
	DeleteByID allows us to delete from any table
	given an id and a statement string

	returns true on success, false otherwise
*/
func DeleteByID(statement string, id int64) bool {
	stmt, err := db.Prepare(statement)
	CheckErr(err)

	res, err := stmt.Exec(id)
	CheckErr(err)

	_, err = res.RowsAffected()
	CheckErr(err)

	return true
}
