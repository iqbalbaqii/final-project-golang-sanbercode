-- +migrate Up
-- +migrate StatementBegin
DROP TABLE public.topik;
CREATE TABLE public.topik (
	id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	judul varchar(256) NOT NULL,
	pertanyaan varchar(256) NULL,
	periode varchar(5) NOT NULL,
	start_date date NOT NULL,
	end_date date NOT NULL,
	created_by varchar(100) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	deleted int4 NULL DEFAULT 0,
	is_release int4 NOT NULL DEFAULT 0,
	target varchar(50) NOT NULL DEFAULT 'ALL'::character varying,
	CONSTRAINT topik_pk PRIMARY KEY (id)
);
-- +migrate StatementEnd