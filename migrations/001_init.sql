
CREATE TABLE IF NOT EXISTS events (
                                      id SERIAL PRIMARY KEY,
                                      user_id INT NOT NULL,
                                      action TEXT NOT NULL,
                                      metadata JSONB,
                                      created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_events_user_id ON events(user_id);
CREATE INDEX IF NOT EXISTS idx_events_created_at ON events(created_at);

CREATE TABLE IF NOT EXISTS aggregated_stats (
                                                id SERIAL PRIMARY KEY,
                                                user_id INT NOT NULL,
                                                event_count INT NOT NULL,
                                                period_start TIMESTAMP NOT NULL,
                                                period_end TIMESTAMP NOT NULL,
                                                created_at TIMESTAMP DEFAULT NOW()
);