package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/royansyahali/blog/payloads/request"
	updaterequest "github.com/royansyahali/blog/payloads/request/updateRequest"
	"github.com/royansyahali/blog/payloads/response"
	"github.com/royansyahali/blog/repositories"
)

type PostRepositoryImp struct {
}

func NewPostRepository() repositories.PostRepository {
	return &PostRepositoryImp{}
}

func (p *PostRepositoryImp) InsertPost(ctx context.Context, tx *sql.Tx, reqPost *request.PostRequest) error {
	repoUser := NewUserRepository()
	user, err := repoUser.FindByIdUser(ctx, tx, reqPost.UserId)
	if err != nil {
		return err
	}
	if user.Role == "admin" {
		query := ("INSERT INTO tb_posts (user_id, title, content, created_at, updated_at) VALUES (?,?,?,?,?)")
		result, err := tx.ExecContext(ctx, query,
			reqPost.UserId,
			reqPost.Title,
			reqPost.Content,
			reqPost.CreateAt,
			reqPost.UpdateAt,
		)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		pc := &request.PostHasCategoryRequest{
			PostId:     int(id),
			CategoryId: reqPost.CategoryId,
		}
		err = pc.Prepare()
		if err != nil {
			return err
		}
		repoPc := NewPostHasCategoryRepository()
		err = repoPc.InsertPostHasCategory(ctx, tx, pc)
		if err != nil {
			return err
		}
		pt := &request.PostHasTagRequest{
			PostId: int(id),
			TagId:  reqPost.TagId,
		}
		err = pt.Prepare()
		if err != nil {
			return err
		}
		repoPt := NewPostHasTagRepository()
		err = repoPt.InsertPostHasTag(ctx, tx, pt)
		if err != nil {
			return err
		}

		return nil
	}
	return errors.New("don't have permission to create this post")

}

func (p *PostRepositoryImp) UpdatePost(ctx context.Context, tx *sql.Tx, reqPost *updaterequest.PostUpdate, id int) error {
	repoUser := NewUserRepository()
	user, err := repoUser.FindByIdUser(ctx, tx, reqPost.UserId)
	if err != nil {
		return err
	}
	if user.Role == "admin" {
		post, err := p.FindByIdPost(ctx, tx, id)
		if err != nil {
			return err
		}
		if reqPost.Title == "" {
			reqPost.Title = post.Title
		}
		if reqPost.Content == "" {
			reqPost.Title = post.Title
		}
		query := ("UPDATE tb_posts SET title = ?, user_id = ?, content = ? , updated_at = ? WHERE id = ?")
		_, err = tx.ExecContext(ctx, query,
			reqPost.Title,
			reqPost.UserId,
			reqPost.Content,
			reqPost.UpdateAt,
			id,
		)
		if err != nil {
			return err
		}
		if reqPost.CategoryId != "" {
			pc := &request.PostHasCategoryRequest{
				PostId:     int(id),
				CategoryId: reqPost.CategoryId,
			}
			err = pc.Prepare()
			if err != nil {
				return err
			}
			repoPc := NewPostHasCategoryRepository()
			err = repoPc.UpdatePostHasCategory(ctx, tx, pc)
			if err != nil {
				return err
			}
		}
		if reqPost.TagId != "" {
			pt := &request.PostHasTagRequest{
				PostId: int(id),
				TagId:  reqPost.TagId,
			}
			err = pt.Prepare()
			if err != nil {
				return err
			}

			repoPt := NewPostHasTagRepository()
			err = repoPt.UpdatePostHasTag(ctx, tx, pt)
			if err != nil {
				return err
			}
		}

		return nil
	}
	return errors.New("don't have permission to update this post")
}

func (p *PostRepositoryImp) DeletePost(ctx context.Context, tx *sql.Tx, id_user, id int) error {
	repoUser := NewUserRepository()
	user, err := repoUser.FindByIdUser(ctx, tx, id_user)
	if err != nil {
		return err
	}
	if user.Role == "admin" {
		_, err := p.FindByIdPost(ctx, tx, id)
		if err != nil {
			return err
		}
		query := ("DELETE FROM tb_posts WHERE id = ?")
		_, err = tx.ExecContext(ctx, query,
			id,
		)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("don't have permission to delete this post")

}

func (p *PostRepositoryImp) FindByIdPost(ctx context.Context, tx *sql.Tx, id int) (response.PostResponse, error) {
	var post response.PostResponse
	query := (`SELECT p.id, p.title, p.content, p.created_at, p.updated_at, t.name, c.name, u.name FROM tb_posts p JOIN tb_post_has_tag pt ON p.id = pt.post_id 
	JOIN tb_tags t ON pt.tag_id = t.id 
	JOIN tb_post_has_category pc ON p.id = pc.post_id 
	JOIN tb_categories c ON pc.category_id = c.id
	JOIN tb_users u ON p.user_id = u.id WHERE p.id = ?
	`)
	result, err := tx.QueryContext(ctx, query,
		id,
	)
	if err != nil {
		return post, err
	}
	defer result.Close()

	var tag, tmpTag string
	var cate, tmp string
	count := -1

	for result.Next() {
		err := result.Scan(&post.Id,
			&post.Title,
			&post.Content,
			&post.CreateAt,
			&post.UpdateAt,
			&tag,
			&cate,
			&post.Name,
		)
		if err != nil {
			return post, err
		}
		if count < 0 {
			tmp += cate
			tmpTag += tag
		} else {
			if !strings.Contains(tmp, cate) {
				tmp += fmt.Sprintf(",%s", cate)
			}
			if !strings.Contains(tmpTag, tag) {
				tmpTag += fmt.Sprintf(",%s", tag)
			}
		}
		count++
	}
	if post.Id != 0 {
		post.Category = strings.Split(tmp, ",")
		post.Tag = strings.Split(tmpTag, ",")
		return post, nil
	}
	return post, errors.New("post is not found in database")

}

func (p *PostRepositoryImp) GetAllPost(ctx context.Context, tx *sql.Tx) ([]response.PostResponse, error) {
	var posts []response.PostResponse
	var post response.PostResponse

	query := (`SELECT p.id, p.title, p.content, p.created_at, p.updated_at, t.name, c.name, u.name FROM tb_posts p JOIN tb_post_has_tag pt ON p.id = pt.post_id 
	JOIN tb_tags t ON pt.tag_id = t.id 
	JOIN tb_post_has_category pc ON p.id = pc.post_id 
	JOIN tb_categories c ON pc.category_id = c.id
	JOIN tb_users u ON p.user_id = u.id`)
	result, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var tag, tmpTag string
	var cate, tmp string
	count := -1
	// var index uint

	for result.Next() {
		err := result.Scan(&post.Id,
			&post.Title,
			&post.Content,
			&post.CreateAt,
			&post.UpdateAt,
			&tag,
			&cate,
			&post.Name,
		)
		if err != nil {
			return posts, err
		}
		if count < 0 {
			tmp += cate
			tmpTag += tag
			post.Category = append(post.Category, cate)
			post.Tag = append(post.Tag, tag)
			posts = append(posts, post)
			count++
		} else if posts[count].Id != post.Id {
			tmp = cate
			tmpTag = tag
			post.Category = []string{}
			post.Tag = []string{}
			post.Category = append(post.Category, cate)
			post.Tag = append(post.Tag, tag)
			posts = append(posts, post)
			count++
		} else {
			if !strings.Contains(tmp, cate) {
				tmp += fmt.Sprintf(",%s", cate)
				posts[count].Category = append(post.Category, cate)
			}
			if !strings.Contains(tmpTag, tag) {
				tmpTag += fmt.Sprintf(",%s", tag)
				posts[count].Tag = append(post.Tag, tag)
			}
		}
	}
	if len(posts) != 0 {
		return posts, nil
	}
	return posts, errors.New("post is empty")

}
