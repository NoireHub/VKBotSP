CREATE TABLE quotes (
   id  SERIAL PRIMARY KEY,
   quote_text varchar not null unique,
   peer_id integer not null
);