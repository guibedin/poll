package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/guibedin/poll/web/domain"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "poll"
	password = "pollpass"
	dbname   = "poll"
)

// sqlRepository implements Repostitory
type sqlRepository struct {
	db *sql.DB
	mq *amqp.Connection
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
	log.Println("Server connected to PostgreSQL")

	// Get MQ connection
	mq, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	log.Println("Server connected to MQ")

	return &sqlRepository{db, mq}
}

// GetPoll returns a single Poll from the database based on its ID
func (r *sqlRepository) GetPoll(id int) (domain.Poll, error) {

	// Get specific poll
	stmt := `SELECT poll_id, title, is_active, is_multiple_choice, created_at, updated_at FROM polls WHERE poll_id = $1;`

	// Query only 1 row with the Poll
	var poll domain.Poll
	err := r.db.QueryRow(stmt, id).Scan(&poll.ID,
		&poll.Title,
		&poll.IsActive,
		&poll.IsMultipleChoice,
		&poll.CreatedAt,
		&poll.UpdatedAt)
	if err != nil {
		panic(err)
	}

	// Get options for Poll
	options, err := r.GetOptionsByPollID(poll.ID)
	if err != nil {
		panic(err)
	}
	poll.Options = options

	// Return poll
	return poll, nil
}

// GetPolls returns all Polls from the database
func (r *sqlRepository) GetPolls() ([]domain.Poll, error) {
	var polls []domain.Poll

	// Get all polls
	stmt := `SELECT poll_id, title, is_active, is_multiple_choice, created_at, updated_at FROM polls;`
	rows, err := r.db.Query(stmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		// Get poll
		var poll domain.Poll
		err = rows.Scan(&poll.ID,
			&poll.Title,
			&poll.IsActive,
			&poll.IsMultipleChoice,
			&poll.CreatedAt,
			&poll.UpdatedAt)
		if err != nil {
			panic(err)
		}

		options, err := r.GetOptionsByPollID(poll.ID)
		if err != nil {
			panic(err)
		}
		poll.Options = options
		polls = append(polls, poll)
	}

	//Return list of polls
	return polls, nil
}

// AddPoll adds the new poll to the database
func (r *sqlRepository) AddPoll(p domain.Poll) int {
	// Save Poll
	stmt := `INSERT INTO polls (title, is_active, is_multiple_choice) VALUES ($1, $2, $3) RETURNING poll_id`

	var id int
	err := r.db.QueryRow(stmt, p.Title, p.IsActive, p.IsMultipleChoice).Scan(&id)
	if err != nil {
		panic(err)
	}

	// Save options to DB
	for _, o := range p.Options {
		r.AddOption(o, id)
	}

	// Return id of new Poll
	return id
}

// GetOptionsByPollID returns all options from a given Poll
func (r *sqlRepository) GetOptionsByPollID(id int) ([]domain.Option, error) {
	var options []domain.Option

	stmt := `SELECT option_id, title, created_at, updated_at FROM options WHERE poll_id = $1;`

	// Get options for poll
	rows, err := r.db.Query(stmt, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var option domain.Option
		err = rows.Scan(&option.ID,
			&option.Title,
			&option.CreatedAt,
			&option.UpdatedAt)
		if err != nil {
			panic(err)
		}

		option.Votes, err = r.GetVoteCountByOptionID(option.ID)
		if err != nil {
			panic(err)
		}
		options = append(options, option)
	}
	return options, nil
}

// AddOption adds an option to the database
func (r *sqlRepository) AddOption(o domain.Option, pollId int) int {
	// Save Options
	stmt := `INSERT INTO options (poll_id, title) VALUES ($1, $2) RETURNING option_id`

	var id int
	err := r.db.QueryRow(stmt, pollId, o.Title).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

// GetVoteCountByOptionID returns the number of votes of a given option
func (r *sqlRepository) GetVoteCountByOptionID(id int) (int, error) {
	// Get vote count
	stmt := `SELECT COUNT(*) FROM votes WHERE option_id = $1;`

	// Get votes
	var votes int
	votesRow := r.db.QueryRow(stmt, id)
	err := votesRow.Scan(&votes)
	if err != nil {
		panic(err)
	}

	return votes, nil
}

func (r *sqlRepository) GetVotesByOptionID(id int) ([]domain.Vote, error) {
	return []domain.Vote{}, nil
}

// AddVote publishes the vote to the queue
func (r *sqlRepository) AddVote(v domain.Vote) error {
	ch, err := r.mq.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"votes", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		panic(err)
	}

	body, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		panic(err)
	}

	return nil
}

// AddVotes publishes the votes to the queue
func (r *sqlRepository) AddVotes(v []domain.Vote) error {
	for _, vote := range v {
		err := r.AddVote(vote)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
