package controller

import (
	"github.com/gin-gonic/gin"
)

// ListJouranls Получить все журналы
// @Summary Список журналов
// @Description Получение списка журналов
// @Tags Journal
// @Accept  json
// @Produce  json
// @Success 200 {array} httputils.Journal
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /journals [get]
func (c *Controller) ListJouranls(ctx *gin.Context) {
}

// ShowJournal Получение списка журналов
// @Summary Список журналов
// @Description Получение списка журналов
// @Tags Journal
// @Accept  json
// @Produce  json
// @Param journal_id path string true "Journal id"
// @Success 200 {object} httputils.ListJournalItems
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /journals/{journal_id} [get]
func (c *Controller) ShowJournal(ctx *gin.Context) {
}

// ListItemsInJournal Получит все объекты журнала
// @Summary Список объектов
// @Description Получение списка объектов с которыми взамодействует журнал, доступных для добавления (используется только для журналов с флагом addition=1)
// @Tags Journal
// @Accept  json
// @Produce  json
// @Param journal_id path string true "Journal id"
// @Success 200 {array} httputils.Item
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /journals/{journal_id}/items [get]
func (c *Controller) ListItemsInJournal(ctx *gin.Context) {
}

// AddItemToJournal Добавить объект
// @Summary Добавить объект
// @Description Добавление позиции в журнал (используется только для журналов с флагом addition=1).
// @Tags Journal
// @Accept  json
// @Produce  json
// @Param item body httputils.Item true "Item"
// @Param journal_id path string true "Journal id"
// @Success 200 {object} httputils.Item
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /journals/{journal_id}/items [post]
func (c *Controller) AddItemToJournal(ctx *gin.Context) {
}

// ShowItemInJournal Получить объект
// @Summary Получить объект
// @Description Получение определенного объекта из журнала.
// @Tags Journal
// @Accept  json
// @Produce  json
// @Param item_id path string true "Item id"
// @Param journal_id path string true "Journal id"
// @Success 200 {object} httputils.JournalItem
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /journals/{journal_id}/items/{item_id} [get]
func (c *Controller) ShowItemInJournal(ctx *gin.Context) {
}

// SaveItemInJournal Сохранение журнала
// @Summary Сохранить журнал
// @Description Сохранение внесенной информации об определенном объекте в журнал.
// @Tags Journal
// @Accept  json
// @Produce  json
// @Param item body array true "Item"
// @Param item_id path string true "Item id"
// @Param journal_id path string true "Journal id"
// @Success 200 {integer} int "Было ли завершено заполнение позиции сегодня -1 ─ была коректировка"
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /journals/{journal_id}/items/{item_id} [put]
func (c *Controller) SaveItemInJournal(ctx *gin.Context) {
}

// DeleteItemFromJournal Удаление объекта
// @Summary Удлить объект
// @Description Удаление итема из журнала. Удаление доступно только для итемов журнала, у которого есть возможность добавления.
// @Tags Journal
// @Accept  json
// @Produce  json
// @Param item_id path string true "Item id"
// @Param journal_id path string true "Journal id"
// @Success 200 {object} httputils.Item
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /journals/{journal_id}/items/{item_id} [delete]
func (c *Controller) DeleteItemFromJournal(ctx *gin.Context) {
}

// CloseJournal Добавить роспись
// @Summary Добавление росписи
// @Description Добавление росписи контролера для закрытия журнала на день. Данная функция доступна только для ежедневных журналов (daily == 0). Роспись - это файл, в формате png размером 250х125.
// @Tags Journal
// @Accept  json
// @Produce  json
// @Param journal_id path string true "Journal id"
// @Success 200 {string} string "answer"
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /journals/{journal_id}/signature [put]
func (c *Controller) CloseJournal(ctx *gin.Context) {
}
