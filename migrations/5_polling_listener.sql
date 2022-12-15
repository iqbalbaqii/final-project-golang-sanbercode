-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE public.listener (
	id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	username varchar NOT NULL,
	target varchar NOT NULL,
	CONSTRAINT listener_pk PRIMARY KEY (username, target),
	CONSTRAINT listener_fk FOREIGN KEY (target) REFERENCES public.target(target)
);
-- +migrate StatementEnd