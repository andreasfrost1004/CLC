CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE guilds (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE guild_memberships (
    id SERIAL PRIMARY KEY,
    guild_id INT NOT NULL REFERENCES guilds(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role TEXT NOT NULL,
    UNIQUE(guild_id, user_id)
);

CREATE TABLE characters (
    id SERIAL PRIMARY KEY,
    guild_membership_id INT NOT NULL REFERENCES guild_memberships(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    class TEXT NOT NULL,
    spec TEXT NOT NULL,
    role TEXT NOT NULL
);

CREATE TABLE raids (
    id SERIAL PRIMARY KEY,
    guild_id INT NOT NULL REFERENCES guilds(id),
    name TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE loot_events (
    id SERIAL PRIMARY KEY,
    raid_id INT NOT NULL REFERENCES raids(id) ON DELETE CASCADE,
    item_id INT NOT NULL,
    player_name TEXT,
    source_boss TEXT,
    timestamp TIMESTAMP,
    raw_row_data JSONB,
    event_index INT NOT NULL
);

CREATE TABLE assignments (
    id SERIAL PRIMARY KEY,
    raid_id INT NOT NULL REFERENCES raids(id) ON DELETE CASCADE,
    version INT NOT NULL,
    item_id INT NOT NULL,
    assigned_character_id INT REFERENCES characters(id),
    reason TEXT,
    is_override BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE item_metadata_cache (
    item_id INT PRIMARY KEY,
    name TEXT,
    icon TEXT,
    slot TEXT,
    tier TEXT,
    weight FLOAT,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
