package store

import "database/sql"

//Post is a posty thingy
type Post struct {
	ID          int    `db:"id"`
	Type        string `db:"post_type"`
	Description string `db:"post_description"`
}

//GetPost gets a post from the store by the ID
func GetPostByID(db *sql.DB, id int) (p Post, err error) {
	query := `
		SELECT
			p.id,
			p.post_type,
			p.post_description
		FROM
			posts
		WHERE
			id=?	
	`

	err = db.QueryRow(query, id).Scan(&p.ID, &p.Type, &p.Description)
	if err != nil {
		return p, err
	}
	return p, nil
}
