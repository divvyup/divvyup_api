package db

import "log"

func IsMemberOf(uid int64) []int64 {
	gids := []int64{}

	rows, err := db.Query("select groupid from membership where userid = ?", uid)
	// If for some reason there is an error
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64

		if err := rows.Scan(&id); err != nil {
			// Something has gone wrong
			log.Fatal(err)
		}

		gids = append(gids, id)
	}

	return gids
}
