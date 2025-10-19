package models

import (
	"time"
)

const UserTokenDataName = "token pengguna"

type GetUserToken struct {
	Id         string    `db:"id"`
	UserId     string    `db:"user_id"`
	Platform   string    `db:"platform"`
	FcmToken   string    `db:"fcm_token"`
	FcmIsValid bool      `db:"fcm_is_valid"`
	ExpiryTime time.Time `db:"expiry_time"`
}

type UpsertUserToken struct {
	UserId     string    `db:"user_id"`
	Platform   string    `db:"platform"`
	FcmToken   string    `db:"fcm_token"`
	ExpiryTime time.Time `db:"expiry_time"`
}

func (GetUserToken) ColumnQuery() string {
	return `
		ut.id,
		ut.user_id,
		ut.platform,
		ut.fcm_token,
		ut.fcm_is_valid,
		ut.expiry_time
	`
}

func (GetUserToken) TableQuery() string {
	return `
		FROM user_tokens ut
	`
}
