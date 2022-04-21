package voting

import (
	"time"
)

type SqlPoll struct {
	PollId    int       `json:"poll_id"`
	Title     string    `json:"title"`
	IsActive  bool      `json:"is_active"`
	CreatedOn time.Time `json:"created_on"`
}

type SqlOption struct {
	OptionId int    `json:"option_id"`
	PollId   int    `json:"poll_id"`
	Title    string `json:"title"`
	Votes    int    `json:"votes"`
}

type JsonPoll struct {
	Title   string       `json:"title"`
	Options []JsonOption `json:"options"`
}

type CompletePoll struct {
	Poll    SqlPoll     `json:"poll"`
	Options []SqlOption `json:"options"`
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

func Get(id int) (CompletePoll, error) {
	// Connect to DB
	db := Connect()
	defer db.Close()

	// Get specific poll
	pollStmt := `SELECT * FROM polls WHERE poll_id = $1`
	var sqlPoll SqlPoll

	optionsStmt := `SELECT * FROM options WHERE poll_id = $1`
	var sqlOptions []SqlOption

	// Query only 1 row with the Poll
	err := db.QueryRow(pollStmt, id).Scan(&sqlPoll.PollId,
		&sqlPoll.Title,
		&sqlPoll.IsActive,
		&sqlPoll.CreatedOn)
	if err != nil {
		panic(err)
	}

	// Query all Options associated with that Poll ID
	rows, err := db.Query(optionsStmt, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Scan through all options to create a list
	for rows.Next() {
		var sqlOption SqlOption
		err = rows.Scan(&sqlOption.OptionId,
			&sqlOption.PollId,
			&sqlOption.Title,
			&sqlOption.Votes)
		if err != nil {
			panic(err)
		}
		sqlOptions = append(sqlOptions, sqlOption)
	}

	// Creates the CompletePoll object that will be sent to the client
	completePoll := CompletePoll{
		sqlPoll,
		sqlOptions,
	}

	// Return poll
	return completePoll, nil
}

func GetAll() ([]CompletePoll, error) {
	// Connect to DB
	db := Connect()
	defer db.Close()

	// Get all polls
	var sqlPolls []SqlPoll
	var completePolls []CompletePoll
	pollStmt := `SELECT * FROM polls`
	optionsStmt := `SELECT * FROM options WHERE poll_id = $1`

	rows, err := db.Query(pollStmt)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var sqlPoll SqlPoll
		err = rows.Scan(&sqlPoll.PollId,
			&sqlPoll.Title,
			&sqlPoll.IsActive,
			&sqlPoll.CreatedOn)
		if err != nil {
			panic(err)
		}
		sqlPolls = append(sqlPolls, sqlPoll)
	}

	for _, sp := range sqlPolls {
		var sqlOptions []SqlOption
		rows, err = db.Query(optionsStmt, sp.PollId)
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var sqlOption SqlOption
			err = rows.Scan(&sqlOption.OptionId,
				&sqlOption.PollId,
				&sqlOption.Title,
				&sqlOption.Votes)
			if err != nil {
				panic(err)
			}
			sqlOptions = append(sqlOptions, sqlOption)
		}
		completePolls = append(completePolls, CompletePoll{
			sp,
			sqlOptions,
		})
	}

	//Return list of polls
	return completePolls, nil
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
