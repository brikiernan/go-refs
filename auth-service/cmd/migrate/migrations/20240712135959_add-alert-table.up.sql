CREATE TABLE IF NOT EXISTS alert (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  contact_id uuid NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),

  CONSTRAINT alert_pkey PRIMARY KEY (id),
  CONSTRAINT alert_contact_id_fkey FOREIGN KEY (contact_id)
    REFERENCES contact (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
);