CREATE TABLE postgres.public.issues (
    id SERIAL PRIMARY KEY,  -- Automatically increments for each new row
    issue_key TEXT UNIQUE NOT NULL,  -- The GHO-1 pattern will be generated with a function
    createdts TIMESTAMP DEFAULT NOW(),  -- Automatically set to current timestamp
    type INT,  -- Will become a foreign key to another table
    title TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE OR REPLACE FUNCTION generate_issue_key() RETURNS TRIGGER AS $$
DECLARE
    issue_count INT;
BEGIN
    -- Count the number of issues to determine the next number in the sequence
    SELECT COUNT(*) INTO issue_count FROM issues;
    
    -- Generate the issue_key as 'GHO-' followed by the next number
    NEW.issue_key := 'GHO-' || (issue_count + 1);

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER issue_key_trigger
BEFORE INSERT ON issues
FOR EACH ROW
EXECUTE FUNCTION generate_issue_key();