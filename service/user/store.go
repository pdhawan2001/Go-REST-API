package user

import (
	"database/sql"
	"fmt"

	"github.com/pdhawan2001/Go-REST-API/types"
)

// this file will help us fetch the data easily, we don't have to write the queries again and again
// we will just write the query functions and we will reuse them, wherever necessary

type Store struct {
	db *sql.DB
}

// Creating a new instance of the Store
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)

	// something went wrong (DB error, no user found, etc.
	// We cannot return a valid user, so we return nil for the first value
	// We return err as the second value so the caller knows what happened
	if err != nil {
		return nil, err
	}

	// What new(types.User) means
	// Allocates memory for types.User
	// Initializes it with zero values
	// Returns a pointer
	// u       // is a *types.User
	// *u      // is the actual User value
	// Equivalent forms:
	// u := new(types.User)
	// and
	// u := &types.User{}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}

		// if UID is 0, that means user not found
		if u.ID == 0 {
			return nil, fmt.Errorf("User not found")
		}
	}

	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
