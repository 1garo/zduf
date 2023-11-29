package teacher

import (
	"context"

	"github.com/1garo/zduf/internal"
	"github.com/jackc/pgx/v5/pgxpool"
)


type Repository struct {
	Conn *pgxpool.Pool
}


// Create insert teacher into the database and return it's `id`
func (r *Repository) Create(t *internal.Teacher) error {
	var id uint

	err := r.Conn.QueryRow(
		context.Background(),
		"insert into teacher(first_name, last_name, age) values ($1, $2, $3) returning id",
		t.FirstName,
		t.LastName,
		t.Age).Scan(&id)

	t.ID = id

	return err
}
