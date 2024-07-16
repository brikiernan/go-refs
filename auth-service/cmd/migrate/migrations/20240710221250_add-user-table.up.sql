CREATE TABLE IF NOT EXISTS "user" (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  first varchar(50) NOT NULL,
  last varchar(50) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  role varchar(50) NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),

  CONSTRAINT user_pkey PRIMARY KEY (id),
  CONSTRAINT user_email_key UNIQUE (email)
);