package store

import "database/sql"

//PostStore is manages Posts in the store
type PostStore struct {
	DB *sql.DB
}

//PostStoreInit returns a Post with the database initialized
func PostStoreInit(db *sql.DB) PostStore {
	return PostStore{DB: db}
}

//Post is a posty thingy
type Post struct {
	ID          int    `db:"id"`
	Type        string `db:"post_type"`
	Description string `db:"post_description"`
}

//Get updates a post from the store by the ID
func (ps *PostStore) GetByID(id int) (p Post, err error) {
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

	err = ps.DB.QueryRow(query, p.ID).Scan(&p.ID, &p.Type, &p.Description)
	if err != nil {
		return p, err
	}
	return p, nil
}
