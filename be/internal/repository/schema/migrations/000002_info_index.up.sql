CREATE EXTENSION pg_trgm;

CREATE INDEX trgm_personal_info_first_name_idx ON personal_info USING gist (first_name gist_trgm_ops);
CREATE INDEX trgm_personal_info_second_name_idx ON personal_info USING gist (second_name gist_trgm_ops);