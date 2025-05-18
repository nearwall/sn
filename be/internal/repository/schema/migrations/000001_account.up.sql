CREATE TABLE account (
    id UUID PRIMARY KEY,
    hashed_pwd TEXT NOT NULL,
    -- 0 - bit pepper on(1)/ off(0)
    -- 1..3 - hash algorithm
    --        0 - argon2
    --        7 -  (only for tests)
    -- 4..15 - pepper version
    hash_features INT2 NOT NULL,
    pwd_updated_at TIMESTAMP NOT NULL,

    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE personal_info (
    account_id UUID PRIMARY KEY REFERENCES account(id) ON DELETE CASCADE,
    first_name  TEXT,
	second_name TEXT,
	birth_date  DATE,
	biography  TEXT,
	city       TEXT,

    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);
