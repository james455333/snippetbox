-- Create a new UTF-8 `snippetbox` database.
CREATE DATABASE snippetbox
    WITH
    ENCODING 'UTF8'
    LC_COLLATE='en_US.utf8'
    LC_CTYPE='en_US.utf8'
    TEMPLATE=template0;

-- Connect to the `snippetbox` database.
\c snippetbox;

-- Create a `snippets` table.
CREATE TABLE snippets (
                          id SERIAL PRIMARY KEY,
                          title VARCHAR(100) NOT NULL,
                          content TEXT NOT NULL,
                          created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          expires TIMESTAMP NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);

-- Add some dummy records.
INSERT INTO snippets (title, content, created, expires) VALUES (
                                                                   'An old silent pond',
                                                                   'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
                                                                   CURRENT_TIMESTAMP,
                                                                   CURRENT_TIMESTAMP + INTERVAL '365 days'
                                                               );

INSERT INTO snippets (title, content, created, expires) VALUES (
                                                                   'Over the wintry forest',
                                                                   'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
                                                                   CURRENT_TIMESTAMP,
                                                                   CURRENT_TIMESTAMP + INTERVAL '365 days'
                                                               );

INSERT INTO snippets (title, content, created, expires) VALUES (
                                                                   'First autumn morning',
                                                                   'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
                                                                   CURRENT_TIMESTAMP,
                                                                   CURRENT_TIMESTAMP + INTERVAL '7 days'
                                                               );

-- Create a new user and grant privileges.
CREATE USER web WITH PASSWORD 'your_password';

-- Grant permissions to the `web` user.
GRANT SELECT, INSERT, UPDATE, DELETE ON snippets TO web;
