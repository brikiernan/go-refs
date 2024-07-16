CREATE TABLE IF NOT EXISTS contact (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),

  CONSTRAINT contact_pkey PRIMARY KEY (id)
);