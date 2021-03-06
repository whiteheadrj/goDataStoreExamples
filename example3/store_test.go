package store

import (
	"database/sql"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPostInit(t *testing.T) {
	Convey("PostInit should return a Post with the database populated", t, func() {

		db, _ := getMockDB(t)
		defer db.Close()

		p := PostInit(db)

		check := Post{
			DB: db,
		}
		So(p, ShouldResemble, check)
	})

}

func TestGetPost(t *testing.T) {

	Convey("Given a valid post id should get post", t, func() {
		db, mock := getMockDB(t)
		defer db.Close()

		columns := getColumns()
		results := sqlmock.NewRows(columns)
		results.AddRow(
			"question",
			"What is the meaning of life?",
		)
		mock.ExpectQuery("SELECT").WillReturnRows(results)

		p := PostInit(db)

		p.ID = 111
		err := p.Get()
		So(err, ShouldBeNil)

		check := Post{
			DB:          db,
			ID:          111,
			Type:        "question",
			Description: "What is the meaning of life?",
		}
		So(p, ShouldResemble, check)

		err = mock.ExpectationsWereMet()
		So(err, ShouldBeNil)
	})
	Convey("Given a invalid id should get error and empty post", t, func() {

		db, mock := getMockDB(t)
		defer db.Close()

		columns := getColumns()
		results := sqlmock.NewRows(columns)
		mock.ExpectQuery("SELECT").WillReturnRows(results)

		p := PostInit(db)

		p.ID = 111
		err := p.Get()

		checkError := sql.ErrNoRows
		So(err, ShouldResemble, checkError)

		check := Post{DB: db, ID: 111}
		So(p, ShouldResemble, check)

		err = mock.ExpectationsWereMet()
		So(err, ShouldBeNil)

	})
}

func getColumns() []string {
	return []string{
		"post_type",
		"post_description",
	}
}

func getMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error initizing the mock DB")
	}

	return mockDB, mock

}
