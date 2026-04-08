CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     name TEXT NOT NULL
);

INSERT INTO users (id, name) VALUES
                                 (1, 'Андрій'),
                                 (2, 'Олександр'),
                                 (3, 'Микола'),
                                 (4, 'Софія'),
                                 (5, 'Дмитро')
ON CONFLICT (id) DO NOTHING;