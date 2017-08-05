package store

import "database/sql"

//Post is a posty thingy
type Post struct {
	DB          *sql.DB
	ID          int    `db:"id"`
	Type        string `db:"post_type"`
	Description string `db:"post_description"`
}

//PostInit returns a Post with the database initialized
func PostInit(db *sql.DB) Post {
	return Post{DB: db}
}

//Get updates a post from the store by the ID
func (p *Post) Get() (err error) {
	query := `
		SELECT
			p.post_type,
			p.post_description
		FROM
			posts
		WHERE
			id=?
	`

	err = p.DB.QueryRow(query, p.ID).Scan(&p.Type, &p.Description)
	if err != nil {
		return err
	}
	return nil
}
