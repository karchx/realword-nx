CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4(),
    email VARCHAR(100) NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    bio VARCHAR(250),
    image VARCHAR(100),
    password_hash VARCHAR(512) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP ,
    PRIMARY KEY("id")
);
