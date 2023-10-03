package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/guibedin/poll/consumer/domain"
	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "poll"
	password = "pollpass"
	dbname   = "poll"
)

// sqlRepository implements Repostitory
type sqlRepository struct {
	db *sql.DB
}

// Sql creates the sqlRepository, which implements Repostitory, that will be used in the server
func NewSqlRepository() *sqlRepository {
	// Get db connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Consumer connected to PostgreSQL")

	return &sqlRepository{db}
}

// AddVote publishes the vote to the queue
func (r *sqlRepository) AddVote(v domain.Vote) error {
	//Update all options based on Ids with votes++
	stmt := `INSERT INTO votes (option_id, poll_id, voter) VALUES ($1, $2, $3)`
	r.db.QueryRow(stmt, v.OptionId, v.PollId, v.Voter)
	return nil
}
