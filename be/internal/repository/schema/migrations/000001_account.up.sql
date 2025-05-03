CREATE TABLE account (
    id UUID PRIMARY KEY,
    hashed_pwd TEXT NOT NULL,

    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE personal_info (
    account_id UUID PRIMARY KEY REFERENCES account(id) NOT NULL,
    first_name  TEXT,
	second_name TEXT,
	birth_date  DATE,
	biography  TEXT,
	city       TEXT,

    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);
