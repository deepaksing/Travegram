CREATE TABLE IF NOT EXISTS "user" (
    user_id SERIAL PRIMARY KEY,
    created_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_ts BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    country VARCHAR(255),
    password_hash TEXT NOT NULL,
    avatar_url TEXT NOT NULL
);

-- Create the Cost_Items table
CREATE TABLE IF NOT EXISTS cost_items (
    cost_item_id SERIAL PRIMARY KEY,
    description TEXT,
    amount DECIMAL(10, 2)
);

-- Create the Itinerary Nodes table
CREATE TABLE IF NOT EXISTS itinerary_nodes (
    node_id SERIAL PRIMARY KEY,
    from_location VARCHAR(255),
    to_location VARCHAR(255),
    user_id INT REFERENCES "user"(user_id),
    date_from DATE,
    date_to DATE,
    cost_items INT[],
    mode_transportation VARCHAR(255)
);

-- Create the Posts table
CREATE TABLE IF NOT EXISTS posts (
    post_id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    from_country VARCHAR(255),
    to_country VARCHAR(255),
    user_id INT REFERENCES "user"(user_id),
    itinerary_nodes INT[]
);

-- Create the Countries table (Optional)
CREATE TABLE IF NOT EXISTS countries (
    country_id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);
