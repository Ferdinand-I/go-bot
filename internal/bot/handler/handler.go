package handler

import "learning_bot/internal/storage"

type Handler struct {
	UserRepo *storage.UserRepo
}

func New(userRepo *storage.UserRepo) *Handler {
	return &Handler{UserRepo: userRepo}
}
