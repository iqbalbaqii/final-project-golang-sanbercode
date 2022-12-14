-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE public.konten (
	id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
	id_topik int8 NULL,
	judul varchar(256) NULL,
	keterangan varchar(256) NULL,
	image_src varchar(256) NULL,
	deleted int4 NOT NULL DEFAULT 0,
	CONSTRAINT konten_pk PRIMARY KEY (id),
	CONSTRAINT konten_fk FOREIGN KEY (id_topik) REFERENCES public.topik(id) ON DELETE CASCADE
);
-- +migrate StatementEnd