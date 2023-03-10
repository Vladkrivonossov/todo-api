package handler

import (
  "errors"
  "github.com/gin-gonic/gin"
  "net/http"
  "strings"
)

const (
  authorizationToken = "Authorization"
  userCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
  header := c.GetHeader(authorizationToken)
  if header == "" {
    NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
    return
  }

  headerParts := strings.Split(header, " ")
  if len(headerParts) != 2 {
    NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
    return
  }
  userId, err := h.servises.Authorization.ParseToken(headerParts[1])
  if err != nil {
    NewErrorResponse(c, http.StatusUnauthorized, err.Error())
    return
  }

  c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
  id, ok := c.Get(userCtx)
  if !ok {
    NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
    return 0, errors.New("user id not found")
  }
  idInt, ok := id.(int)
  if !ok {
    NewErrorResponse(c, http.StatusInternalServerError, "user id is not valid type")
    return 0, errors.New("user id is not valid type")
  }

  return idInt, nil
}