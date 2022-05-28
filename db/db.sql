CREATE TABLE polls (
    poll_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    is_active BOOLEAN NOT NULL,
    is_multiple_choice BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE options (
    option_id SERIAL PRIMARY KEY,
    poll_id INT NOT NULL,
    title TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (poll_id) REFERENCES polls (poll_id)
);

CREATE TABLE votes (
    vote_id SERIAL PRIMARY KEY,
    option_id INT NOT NULL,
    poll_id INT NOT NULL,
    voter TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (option_id) REFERENCES options (option_id),
    FOREIGN KEY (poll_id) REFERENCES polls (poll_id)
);

/* Create function used by trigger - updated_at */
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

/* Triggers for every table */
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON polls
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON options
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON votes
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();