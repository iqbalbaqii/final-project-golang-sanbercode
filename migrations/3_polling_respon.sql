-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE public.respon (
	id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	id_topik int8 NOT NULL,
	response int8 NULL,
	created_by varchar(10) NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT respon_pk PRIMARY KEY (id),
	CONSTRAINT respon_konten FOREIGN KEY (response) REFERENCES public.konten(id),
	CONSTRAINT respon_topik FOREIGN KEY (id_topik) REFERENCES public.topik(id) ON DELETE CASCADE
);
-- +migrate StatementEnd