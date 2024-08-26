package utils

import (
	"encoding/json"
	"strings"

	"gorm.io/gorm"
)

func Dump(v interface{}) string {
	json, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(json)
}

func BuildSortQuery(db *gorm.DB, sort string) {
	if sort == "" {
		return
	}

	splitted := strings.Split(sort, ",")
	for _, s := range splitted {
		if strings.HasPrefix(s, "-") {
			db.Order(strings.TrimLeft(s, "-") + " DESC")
		} else {
			db.Order(s)
		}
	}
}
