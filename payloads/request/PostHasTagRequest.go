package request

import (
	"errors"
	"strconv"
	"strings"

	"blog/entities"
)

type PostHasTagRequest struct {
	PostId int    `json:"post_id"`
	TagId  string `json:"tag_id"`
	Tag    []int
}

func (p *PostHasTagRequest) Prepare() error {
	p.TagId = strings.ToLower(p.TagId)
	tagStr := strings.Split(p.TagId, ",")
	if tagStr[len(tagStr)-1] == "" || string(tagStr[len(tagStr)-1][0]) == " " {
		tagStr = tagStr[:len(tagStr)-1]
	}
	for _, v := range tagStr {
		tagId, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		p.Tag = append(p.Tag, tagId)
	}
	return nil
}

func (p *PostHasTagRequest) Check(t []entities.Tag) error {
	for _, v := range p.Tag {
		found := false
		for _, x := range t {
			if v == x.Id {
				found = true
				break
			}
		}
		if !found {
			return errors.New("id tag don't exists in database")
		}
	}
	return nil
}
