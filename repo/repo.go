package repo

import (
	"strings"
	"task-app/models"
)

type WhereObj struct {
	field     string
	condition string
	value     string
}

func getPagingInfo(query *models.APIPagingDto, count int) *models.PagingInfo {
	var hasNextPage bool
	next := int64((query.Page * query.Limit) - count)
	if next < 0 && query.Limit > 0 {
		hasNextPage = true
	}
	pagingInfo := &models.PagingInfo{
		HasNextPage: hasNextPage,
		Page:        query.Page,
	}
	return pagingInfo
}

func getWhereObject(filter string) *WhereObj {
	splitted := strings.Split(filter, "|")

	if len(splitted) < 3 {
		return nil
	}

	return &WhereObj{
		field:     splitted[0],
		condition: splitted[1],
		value:     splitted[2],
	}
}
