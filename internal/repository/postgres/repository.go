package postgres

import (
    "database/sql"
    "time"
    "github.com/mfaxmodem/tracker/internal/domain/models"
)

type Repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{
        db: db,
    }
}

func (r *Repository) CreateUser(user *models.User) error {
    query := `
        INSERT INTO users (name, email, password_hash, role, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $5)
        RETURNING id`
    
    now := time.Now()
    return r.db.QueryRow(query, 
        user.Name, 
        user.Email, 
        user.PasswordHash, 
        user.Role,
        now,
    ).Scan(&user.ID)
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
    user := &models.User{}
    query := `
        SELECT id, name, email, password_hash, role, created_at, updated_at
        FROM users 
        WHERE email = $1`
    
    err := r.db.QueryRow(query, email).Scan(
        &user.ID,
        &user.Name,
        &user.Email,
        &user.PasswordHash,
        &user.Role,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (r *Repository) GetAllVisitors() ([]models.User, error) {
    query := `
        SELECT id, name, email, role, created_at, updated_at
        FROM users 
        WHERE role = 'visitor'`
    
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var visitors []models.User
    for rows.Next() {
        var v models.User
        if err := rows.Scan(
            &v.ID,
            &v.Name,
            &v.Email,
            &v.Role,
            &v.CreatedAt,
            &v.UpdatedAt,
        ); err != nil {
            return nil, err
        }
        visitors = append(visitors, v)
    }
    return visitors, nil
}

func (r *Repository) UpdateUser(user *models.User) error {
    query := `
        UPDATE users 
        SET name = $1, email = $2, role = $3, updated_at = $4
        WHERE id = $5`
    
    _, err := r.db.Exec(query,
        user.Name,
        user.Email,
        user.Role,
        time.Now(),
        user.ID,
    )
    return err
}

func (r *Repository) DeleteUser(id int64) error {
    query := `DELETE FROM users WHERE id = $1`
    _, err := r.db.Exec(query, id)
    return err
}

func (r *Repository) SaveStore(store *models.Store) error {
    query := `
        INSERT INTO stores (name, address, latitude, longitude, manager_name, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $6)
        RETURNING id`
    
    now := time.Now()
    return r.db.QueryRow(query,
        store.Name,
        store.Address,
        store.Latitude,
        store.Longitude,
        store.ManagerName,
        now,
    ).Scan(&store.ID)
}

func (r *Repository) GetStores() ([]models.Store, error) {
    query := `
        SELECT id, name, address, latitude, longitude, manager_name, created_at, updated_at
        FROM stores`
    
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var stores []models.Store
    for rows.Next() {
        var s models.Store
        if err := rows.Scan(
            &s.ID,
            &s.Name,
            &s.Address,
            &s.Latitude,
            &s.Longitude,
            &s.ManagerName,
            &s.CreatedAt,
            &s.UpdatedAt,
        ); err != nil {
            return nil, err
        }
        stores = append(stores, s)
    }
    return stores, nil
}

func (r *Repository) UpdateStore(store *models.Store) error {
    query := `
        UPDATE stores 
        SET name = $1, address = $2, latitude = $3, longitude = $4, 
            manager_name = $5, updated_at = $6
        WHERE id = $7`
    
    _, err := r.db.Exec(query,
        store.Name,
        store.Address,
        store.Latitude,
        store.Longitude,
        store.ManagerName,
        time.Now(),
        store.ID,
    )
    return err
}

func (r *Repository) DeleteStore(id int64) error {
    query := `DELETE FROM stores WHERE id = $1`
    _, err := r.db.Exec(query, id)
    return err
}

func (r *Repository) SaveRoute(route *models.Route) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }

    query := `
        INSERT INTO routes (visitor_id, status, start_date, end_date, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $5)
        RETURNING id`
    
    now := time.Now()
    err = tx.QueryRow(query,
        route.VisitorID,
        route.Status,
        route.StartDate,
        route.EndDate,
        now,
    ).Scan(&route.ID)
    if err != nil {
        tx.Rollback()
        return err
    }

    // Insert route stores
    for i, storeID := range route.StoreIDs {
        _, err = tx.Exec(`
            INSERT INTO route_stores (route_id, store_id, visit_order)
            VALUES ($1, $2, $3)`,
            route.ID, storeID, i+1,
        )
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    return tx.Commit()
}

func (r *Repository) GetAllRoutes() ([]models.Route, error) {
    query := `
        SELECT r.id, r.visitor_id, r.status, r.start_date, r.end_date, 
               r.created_at, r.updated_at,
               array_agg(rs.store_id) as store_ids
        FROM routes r
        LEFT JOIN route_stores rs ON r.id = rs.route_id
        GROUP BY r.id`

    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var routes []models.Route
    for rows.Next() {
        var r models.Route
        if err := rows.Scan(
            &r.ID,
            &r.VisitorID,
            &r.Status,
            &r.StartDate,
            &r.EndDate,
            &r.CreatedAt,
            &r.UpdatedAt,
            &r.StoreIDs,
        ); err != nil {
            return nil, err
        }
        routes = append(routes, r)
    }
    return routes, nil
}

func (r *Repository) GetVisitorRoutes(visitorID int64) ([]models.Route, error) {
    query := `
        SELECT r.id, r.visitor_id, r.status, r.start_date, r.end_date, 
               r.created_at, r.updated_at,
               array_agg(rs.store_id) as store_ids
        FROM routes r
        LEFT JOIN route_stores rs ON r.id = rs.route_id
        WHERE r.visitor_id = $1
        GROUP BY r.id`

    rows, err := r.db.Query(query, visitorID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var routes []models.Route
    for rows.Next() {
        var r models.Route
        if err := rows.Scan(
            &r.ID,
            &r.VisitorID,
            &r.Status,
            &r.StartDate,
            &r.EndDate,
            &r.CreatedAt,
            &r.UpdatedAt,
            &r.StoreIDs,
        ); err != nil {
            return nil, err
        }
        routes = append(routes, r)
    }
    return routes, nil
}

func (r *Repository) UpdateRoute(route *models.Route) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }

    query := `
        UPDATE routes 
        SET status = $1, start_date = $2, end_date = $3, updated_at = $4
        WHERE id = $5`

    _, err = tx.Exec(query,
        route.Status,
        route.StartDate,
        route.EndDate,
        time.Now(),
        route.ID,
    )
    if err != nil {
        tx.Rollback()
        return err
    }

    // Update route stores
    _, err = tx.Exec(`DELETE FROM route_stores WHERE route_id = $1`, route.ID)
    if err != nil {
        tx.Rollback()
        return err
    }

    for i, storeID := range route.StoreIDs {
        _, err = tx.Exec(`
            INSERT INTO route_stores (route_id, store_id, visit_order)
            VALUES ($1, $2, $3)`,
            route.ID, storeID, i+1,
        )
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    return tx.Commit()
}

func (r *Repository) DeleteRoute(id int64) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec(`DELETE FROM route_stores WHERE route_id = $1`, id)
    if err != nil {
        tx.Rollback()
        return err
    }

    _, err = tx.Exec(`DELETE FROM routes WHERE id = $1`, id)
    if err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit()
}

func (r *Repository) SaveLocation(location *models.Location) error {
    query := `
        INSERT INTO locations (visitor_id, latitude, longitude, timestamp)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

    return r.db.QueryRow(query,
        location.VisitorID,
        location.Latitude,
        location.Longitude,
        location.Timestamp,
    ).Scan(&location.ID)
}

func (r *Repository) GetLocations(visitorID int64) ([]models.Location, error) {
    query := `
        SELECT id, visitor_id, latitude, longitude, timestamp
        FROM locations
        WHERE visitor_id = $1
        ORDER BY timestamp DESC`

    rows, err := r.db.Query(query, visitorID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var locations []models.Location
    for rows.Next() {
        var l models.Location
        if err := rows.Scan(
            &l.ID,
            &l.VisitorID,
            &l.Latitude,
            &l.Longitude,
            &l.Timestamp,
        ); err != nil {
            return nil, err
        }
        locations = append(locations, l)
    }
    return locations, nil
}