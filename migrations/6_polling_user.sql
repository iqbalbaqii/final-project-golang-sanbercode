-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE public."user" (
	id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	username varchar NOT NULL,
	"password" varchar NOT NULL,
	email varchar NULL,
	full_name varchar NOT NULL,
	"level" int8 NOT NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	gender varchar NULL,
	CONSTRAINT user_pk PRIMARY KEY (username)
);
-- +migrate StatementEnd