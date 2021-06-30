
-- people definition
CREATE TABLE people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT);
CREATE TABLE register_data (id INTEGER PRIMARY KEY,  reg_key TEXT, reg_value TEXT);
CREATE TABLE contacts (
                          contact_id INTEGER PRIMARY KEY,
                          first_name TEXT NOT NULL,
                          last_name TEXT NOT NULL,
                          email TEXT NOT NULL UNIQUE,
                          phone TEXT NOT NULL UNIQUE
);