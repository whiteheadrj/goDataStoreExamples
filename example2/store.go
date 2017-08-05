package store

import "database/sql"

//Post is a posty thingy
type Post struct {
	ID          int    `db:"id"`
	Type        string `db:"post_type"`
	Description string `db:"post_description"`
}

//Get updates a post from the store by the ID
func (p *Post) Get(db *sql.DB) (err error) {
	query := `
		SELECT
			p.post_type,
			p.post_description
		FROM
			posts
		WHERE
			id=?	
	`

	err = db.QueryRow(query, p.ID).Scan(&p.Type, &p.Description)
	if err != nil {
		return err
	}
	return nil
}
