package database

import (
	"fmt"
	"net/http"
	i "token-master/internal"

	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
)

type Token struct {
	JWRToken *jwt.Token
	Claims   *i.Claims
	Raw      string
}

func (u *User) UpdateUserRefreshToken(currentToken string, updatedToken string) *i.Error {
	err := u.checkCurrentUserRefreshTokenIsEqual(u.ID, currentToken)
	if err != nil {
		err.Error = fmt.Sprintf("UpdateUserRefreshToken equal token: %v", err)
		return err
	}

	err = u.insertUserRefreshToken(u.ID, updatedToken)
	if err != nil {
		err.Error = fmt.Sprintf("UpdateUserRefreshToken: %v", err)
		return err
	}

	return nil
}

func (u *User) checkCurrentUserRefreshTokenIsEqual(id int, currentToken string) *i.Error {
	row := i.DB.QueryRow(GetEqualUserRefreshTokenQuery, id, currentToken)
	var refreshToken string
	err := row.Scan(&refreshToken)
	if err != nil {
		return i.DetailError(http.StatusInternalServerError, fmt.Errorf("database: %v", err))
	}

	return nil
}

func (u *User) insertUserRefreshToken(id int, token string) *i.Error {
	if _, err := i.DB.Exec(UpdateUserRefreshTokenQuery, token, id); err != nil {
		return i.DetailError(http.StatusInternalServerError, err)
	}

	return nil
}

func (t *Token) InsertRefreshToken(userID int) *i.Error {
	if _, err := i.DB.Exec(InsertRefreshTokenQuery, userID, t.Raw); err != nil {
		return i.DetailError(http.StatusInternalServerError, fmt.Errorf("InsertRefreshToken database: %v", err))
	}

	return nil
}
