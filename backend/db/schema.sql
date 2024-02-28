CREATE TABLE IF NOT EXISTS rank (
    id INTEGER PRIMARY KEY,
    username TEXT NOT NULL,
    score INTEGER NOT NULL,
    difficulty TEXT CHECK(difficulty in ("easy", "medium", "hard")) NOT NULL,
    created_at TEXT DEFAULT (DATETIME('now'))
)