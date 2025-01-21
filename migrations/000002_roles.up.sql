CREATE TABLE roles (
    role_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,  -- Unique identifier for the role
    role_name VARCHAR(50) NOT NULL UNIQUE,           -- Role name (e.g., admin, user, moderator)
    description VARCHAR(255),                   -- Optional description of the role
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Update timestamp
);
