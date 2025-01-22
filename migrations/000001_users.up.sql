CREATE TABLE users (
    user_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,  -- Unique identifier for the user
    username VARCHAR(50) NOT NULL UNIQUE,          -- Username for login
    email VARCHAR(100) NOT NULL UNIQUE,            -- Email address
    password_hash VARCHAR(255) NOT NULL,           -- Hashed password
    first_name VARCHAR(50),                        -- First name
    last_name VARCHAR(50),                         -- Last name
    role_id INT UNSIGNED NOT NULL,                 -- Foreign key referencing roles table
    is_active BOOLEAN DEFAULT TRUE,                -- To enable/disable accounts
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Account creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Update timestamp
);
