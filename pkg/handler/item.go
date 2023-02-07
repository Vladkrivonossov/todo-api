package handler

import (
  "github.com/gin-gonic/gin"
  todo "go-api"
  "net/http"
  "strconv"
)

func (h *Handler) createItem(c *gin.Context)  {
  userId, err := getUserId(c)
  if err != nil {
    return
  }

  listId, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
    return
  }

  var input todo.TodoItem
  if err := c.BindJSON(&input); err != nil {
    NewErrorResponse(c, http.StatusBadRequest, err.Error())
    return
  }

  id, err := h.servises.TodoItem.Create(userId, listId, input)
  if err != nil {
    NewErrorResponse(c, http.StatusInternalServerError, err.Error())
    return
  }

  c.JSON(http.StatusOK, map[string]interface{}{
    "id": id,
  })
}

func (h *Handler) getAllItems(c *gin.Context)  {
  userId, err := getUserId(c)
  if err != nil {
    return
  }

  listId, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
    return
  }

  items, err := h.servises.TodoItem.GetAll(userId, listId)
  if err != nil {
    NewErrorResponse(c, http.StatusBadRequest, err.Error())
    return
  }

  c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(c *gin.Context)  {
  userId, err := getUserId(c)
  if err != nil {
    return
  }

  itemId, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
    return
  }

  item, err := h.servises.TodoItem.GetById(userId, itemId)
  if err != nil {
    NewErrorResponse(c, http.StatusBadRequest, err.Error())
    return
  }

  c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context)  {
  userId, err := getUserId(c)
  if err != nil {
    return
  }

  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    NewErrorResponse(c, http.StatusBadRequest, err.Error())
    return
  }

  var input todo.UpdateItemInput
  if err := c.BindJSON(&input); err != nil {
    NewErrorResponse(c, http.StatusBadRequest, err.Error())
    return
  }

  if err := h.servises.TodoItem.Update(userId, id, input); err != nil {
    NewErrorResponse(c, http.StatusInternalServerError, err.Error())
    return
  }

  c.JSON(http.StatusOK, statusResponse{"ok"})
}


func (h *Handler) deleteItem(c *gin.Context)  {
  userId, err := getUserId(c)
  if err != nil {
    return
  }

  itemId, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
    return
  }

  err = h.servises.TodoItem.Delete(userId, itemId)
  if err != nil {
    NewErrorResponse(c, http.StatusBadRequest, err.Error())
    return
  }

  c.JSON(http.StatusOK, statusResponse{"ok"})
}