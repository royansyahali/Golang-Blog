package request

import (
	"errors"
	"strconv"
	"strings"

	"github.com/royansyahali/blog/entities"
)

type PostHasCategoryRequest struct {
	PostId     int    `json:"post_id"`
	CategoryId string `json:"category_id"`
	Category   []int
}

func (p *PostHasCategoryRequest) Prepare() error {
	p.CategoryId = strings.ToLower(p.CategoryId)
	tagStr := strings.Split(p.CategoryId, ",")
	if tagStr[len(tagStr)-1] == "" || string(tagStr[len(tagStr)-1][0]) == " " {
		tagStr = tagStr[:len(tagStr)-1]
	}
	for _, v := range tagStr {
		category, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		p.Category = append(p.Category, category)
	}
	return nil
}

func (p *PostHasCategoryRequest) Check(t []entities.Category) error {
	for _, v := range p.Category {
		found := false
		for _, x := range t {
			if v == x.Id {
				found = true
				break
			}
		}
		if !found {
			return errors.New("id category don't exists in database")
		}
	}
	return nil
}
