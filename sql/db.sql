CREATE TABLE IF NOT EXISTS endpoints (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    interval INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS events (
    id TEXT PRIMARY KEY,
    url TEXT NOT NULL,
    status TEXT NOT NULL,
    response_time INTEGER NOT NULL
);