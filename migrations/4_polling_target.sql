-- +migrate Up
-- +migrate StatementBegin
DROP TABLE public.target;
CREATE TABLE public.target (
	id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	target varchar NOT NULL,
	description varchar NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT target_pk PRIMARY KEY (target)
);
-- +migrate StatementEnd