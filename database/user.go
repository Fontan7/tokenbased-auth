package database

import (
	"database/sql"
	"fmt"
	"net/http"
	i "token-master/internal"

	_ "github.com/lib/pq"
)

type User struct {
	ID          int
	UserName    string
	DisplayName string
	Email       string
	SpotifyID   string
	AppleID     string
}

func (u *User) GetUserByEmail(email string) *i.Error {
	//go get this in the db
	if found, err := u.getUserByEmail(email); err != nil || !found {
		err.Error = fmt.Sprintf("GetUserByEmail: no user found with this email or an error occurred: %v", err)
		return err
	}

	return nil
}

func (u *User) getUserByEmail(email string) (bool, *i.Error) {
	row := i.DB.QueryRow(GetUserByEmailQuery, email)
	err := row.Scan(&u.ID, &u.UserName, &u.DisplayName, &u.Email, &u.SpotifyID, &u.AppleID)
	if err != nil {
		if err == sql.ErrNoRows {
			// There were no rows, but otherwise no error occurred
			return false, nil
		}
		return false, i.DetailError(http.StatusInternalServerError, fmt.Errorf("database: %v", err))
	}

	return true, nil
}

func (u *User) GetUserById(id int) *i.Error {
	//control the input
	uid, err := i.IntIdtoUint(id)
	if err != nil {
		return i.DetailError(http.StatusInternalServerError, fmt.Errorf("GetUserById: %v", err))
	}

	if found, err := u.getUserById(uid); err != nil || !found {
		err.Error = fmt.Sprintf("GetUserById: no user found with this id or an error occurred: %v", err)
		return err
	}

	return nil
}

func (u *User) getUserById(id uint) (bool, *i.Error) {
	row := i.DB.QueryRow(GetUserByIdQuery, id)
	err := row.Scan(&u.ID, &u.UserName, &u.DisplayName, &u.Email, &u.SpotifyID, &u.AppleID)
	if err != nil {
		if err == sql.ErrNoRows {
			// There were no rows, but otherwise no error occurred
			return false, nil
		}
		return false, i.DetailError(http.StatusInternalServerError, fmt.Errorf("database: %v", err))
	}

	return true, nil
}

func (u *User) IsLoggedIn(id int) (bool, *i.Error) {
	refreshToken, err := u.GetUserRefreshToken(id)
	if err != nil {
		err.Error = fmt.Sprintf("IsLoggedIn: %v", err)
		return false, err
	}

	// check if the user is already logged in with the token
	if refreshToken == "" || len(refreshToken) < 15 {
		return false, nil
	}

	return true, nil
}

func (u *User) GetUserRefreshToken(id int) (string, *i.Error) {
	uid, err := i.IntIdtoUint(id)
	if err != nil {
		err.Error = fmt.Sprintf("GetUserRefreshToken: %v", err)
		return "", err
	}

	var refreshToken string
	row := i.DB.QueryRow(GetUserRefreshTokenQuery, uid)
	serr := row.Scan(&refreshToken)
	if serr != nil {
		return "", i.DetailError(http.StatusInternalServerError, fmt.Errorf("GetUserRefreshToken database: %v", serr))
	}

	return refreshToken, nil
}
