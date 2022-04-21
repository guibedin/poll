package voting

import (
	"time"
)

type SqlPoll struct {
	PollId    int
	Title     string
	IsActive  bool
	CreatedOn time.Time
}

type SqlOption struct {
	OptionId int
	PollId   int
	Title    string
	Votes    int
}

type JsonPoll struct {
	Title   string       `json:"title"`
	Options []JsonOption `json:"options"`
}

type MqVote struct {
	PollId  int       `json:"poll_id"`
	Options OptionIds `json:"option_ids"`
}

type JsonOption struct {
	Title string `json:"title"`
}

type JsonOptionIds struct {
	OptionIds OptionIds `json:"option_ids"`
}

type OptionIds []int

func (p *JsonPoll) Save() int {
	// Connect to DB
	db := Connect()
	defer db.Close()

	// Save Poll
	pollStmt := `INSERT INTO polls (title, is_active, created_on) VALUES ($1, $2, $3) RETURNING poll_id`
	id := 0
	err := db.QueryRow(pollStmt, p.Title, true, time.Now()).Scan(&id)
	if err != nil {
		panic(err)
	}

	// Save Options
	optionsStmt := `INSERT INTO options (poll_id, title, votes) VALUES ($1, $2, $3)`
	for _, option := range p.Options {
		_, err = db.Exec(optionsStmt, id, option.Title, 0)
		if err != nil {
			panic(err)
		}
	}

	// Return id of new Poll
	return id
}

func Get(id int) (SqlPoll, error) {
	// Connect to DB
	db := Connect()
	defer db.Close()

	// Get specific poll
	pollStmt := `SELECT * FROM polls WHERE poll_id = $1`
	optionsStmt := `SELECT * FROM options WHERE poll_id = $1`

	// Return poll
	return SqlPoll{}, nil
}

func GetAll() ([]SqlPoll, error) {
	// Connect to DB

	// Get all polls

	//Return list of polls
	return nil, nil
}

func Vote(ids OptionIds) {
	// Connect to DB
	db := Connect()
	defer db.Close()

	// Update all options based on Ids with votes++
	selectStmt := `SELECT votes FROM options WHERE option_id = $1 AND poll_id = $2`
	updateStmt := `UPDATE options SET (votes = $1) WHERE option_id = $2 and poll_id = $3`
	var votes int
	for _, id := range ids {
		row := db.QueryRow(selectStmt, id)
		err := row.Scan(&votes)
		if err != nil {
			panic(err)
		} else {
			_, err = db.Exec(updateStmt, votes+1, id)
			if err != nil {
				panic(err)
			}
		}
	}
}
