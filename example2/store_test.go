package store

import (
	"database/sql"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	. "github.com/smartystreets/goconvey/convey"
)

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

		p := Post{ID: 111}
		err := p.Get(db)
		So(err, ShouldBeNil)

		check := Post{
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

		p := Post{ID: 111}
		err := p.Get(db)

		checkError := sql.ErrNoRows
		So(err, ShouldResemble, checkError)

		check := Post{ID: 111}
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
