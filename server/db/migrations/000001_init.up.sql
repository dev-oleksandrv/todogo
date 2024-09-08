CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    content VARCHAR(255) NOT NULL,
    description TEXT,
    is_checked BOOLEAN DEFAULT false,
    position INTEGER DEFAULT 0 NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);