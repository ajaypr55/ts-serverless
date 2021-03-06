package models

import (
	uuid "github.com/satori/go.uuid"
)

type CreateVoteModel struct {
	ObjectId         uuid.UUID `json:"objectId"`
	OwnerUserId      uuid.UUID `json:"ownerUserId"`
	OwnerDisplayName string    `json:"ownerDisplayName"`
	OwnerAvatar      string    `json:"ownerAvatar"`
	PostId           uuid.UUID `json:"postId"`
	TypeId           int       `json:"type"`
	CreatedDate      int64     `json:"created_date"`
}
