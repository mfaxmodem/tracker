package migrations

import (
    "database/sql"
)

func InitDB(db *sql.DB) error {
    // Create users table
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            password_hash VARCHAR(255) NOT NULL,
            role VARCHAR(20) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );

        CREATE TABLE IF NOT EXISTS stores (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            address TEXT NOT NULL,
            latitude DECIMAL(10,8) NOT NULL,
            longitude DECIMAL(11,8) NOT NULL,
            manager_name VARCHAR(100),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );

        CREATE TABLE IF NOT EXISTS routes (
            id SERIAL PRIMARY KEY,
            visitor_id INTEGER REFERENCES users(id),
            status VARCHAR(20) NOT NULL,
            start_date TIMESTAMP NOT NULL,
            end_date TIMESTAMP NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );

        CREATE TABLE IF NOT EXISTS route_stores (
            route_id INTEGER REFERENCES routes(id),
            store_id INTEGER REFERENCES stores(id),
            visit_order INTEGER NOT NULL,
            PRIMARY KEY (route_id, store_id)
        );
    `)
    
    if err != nil {
        return err
    }

    // Insert default admin user
    _, err = db.Exec(`
        INSERT INTO users (name, email, password_hash, role)
        VALUES ('Admin', 'admin@example.com', '$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.X/6wBxoZQwyzpYVQBqwpJf9nrQi8i', 'admin')
        ON CONFLICT (email) DO NOTHING;
    `)

    return err
}