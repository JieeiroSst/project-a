package delivery

import (
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/post/internal/query/usecase"
)

type postHttp struct {
	usecase usecase.PostUsecase
}

type PostHttp interface {
	CreatePosts(posts *model.Posts) error
	UpdatePosts(id int,posts *model.Posts) error
	DeletePosts(id int) error
	PostsById(id int) (*model.Posts, error)
	CreateProfile(profile *model.Profiles) error
	UpdateProfile(id int,profile *model.Profiles) error
	ProfileById(id int) (*model.Profiles, error)
	CreatePostMetas(metas *model.PostMetas) error
	UpdatePostMetas(id int,metas *model.PostMetas) error
	DeletePostMetas(id int) error
	PostMetasById(id int) (*model.PostMetas, error)
	CreateComment(comment *model.PostComments) error
	CreateCategories(categories *model.Categories) error
	UpdateCategories(id int,categories *model.Categories) error
	DeleteCategories(id int) error
	CategoriesById(id int) (*model.Categories, error)
	PostsAll(post *model.Posts, pagination *model.PaginationPage) (*[]model.Posts, error)
	ProfileAll(profile *model.Profiles, pagination *model.PaginationPage) (*[]model.Profiles, error)
	PostMetasAll(postMetas *model.PostMetas, pagination *model.PaginationPage) (*[]model.PostMetas, error)
	CategoriesAll(category *model.Categories, pagination *model.PaginationPage) (*[]model.Categories, error)
	CommentAllPost(idPost int,comment *model.PostComments, pagination *model.PaginationPage) (*[]model.PostComments, error)
	PublishPost(id int) error
	RemoveComment(id int) error
	ListPublishPost(post *model.Posts, pagination *model.PaginationPage) (*[]model.Posts, error)
	ListNotPublishPost(post *model.Posts, pagination *model.PaginationPage) (*[]model.Posts, error)
}

func NewPostHttp(usecase usecase.PostUsecase) PostHttp {
	return &postHttp{
		usecase:usecase,
	}
}

func (p *postHttp) CreatePosts(posts *model.Posts) error {
	if err := p.usecase.CreatePosts(posts); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) UpdatePosts(id int, posts *model.Posts) error {
	if err := p.usecase.UpdatePosts(id, posts); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) DeletePosts(id int) error {
	if err := p.usecase.DeletePosts(id); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) PostsById(id int) (*model.Posts, error) {
	post, err := p.usecase.PostsById(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *postHttp) CreateProfile(profile *model.Profiles) error {
	if err := p.usecase.CreateProfile(profile); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) UpdateProfile(id int, profile *model.Profiles) error {
	if err := p.usecase.UpdateProfile(id, profile); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) ProfileById(id int) (*model.Profiles, error) {
	profile, err := p.usecase.ProfileById(id)
	if err != nil {
		return nil ,err
	}
	return profile, nil
}

func (p *postHttp) CreatePostMetas(metas *model.PostMetas) error {
	if err := p.usecase.CreatePostMetas(metas); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) UpdatePostMetas(id int, metas *model.PostMetas) error {
	if err := p.usecase.UpdatePostMetas(id, metas); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) DeletePostMetas(id int) error {
	if err := p.usecase.DeletePostMetas(id); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) PostMetasById(id int) (*model.PostMetas, error) {
	postMetas, err := p.usecase.PostMetasById(id)
	if err != nil {
		return nil ,err
	}
	return postMetas, nil
}

func (p *postHttp) CreateComment(comment *model.PostComments) error {
	if err := p.usecase.CreateComment(comment); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) CreateCategories(categories *model.Categories) error {
	if err := p.usecase.CreateCategories(categories); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) UpdateCategories(id int, categories *model.Categories) error {
	if err := p.usecase.UpdateCategories(id, categories); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) DeleteCategories(id int) error {
	if err := p.usecase.DeleteCategories(id); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) CategoriesById(id int) (*model.Categories, error) {
	category, err := p.usecase.CategoriesById(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (p *postHttp) PostsAll(post *model.Posts, pagination *model.PaginationPage) (*[]model.Posts, error) {
	posts, err := p.usecase.PostsAll(post, pagination)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postHttp) ProfileAll(profile *model.Profiles, pagination *model.PaginationPage) (*[]model.Profiles, error) {
	profiles, err := p.usecase.ProfileAll(profile, pagination)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (p *postHttp) PostMetasAll(postMetas *model.PostMetas, pagination *model.PaginationPage) (*[]model.PostMetas, error) {
	postMeta, err := p.usecase.PostMetasAll(postMetas, pagination)
	if err != nil {
		return nil ,err
	}
	return postMeta, nil
}

func (p *postHttp) CategoriesAll(category *model.Categories, pagination *model.PaginationPage) (*[]model.Categories, error) {
	categories, err := p.usecase.CategoriesAll(category, pagination)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (p *postHttp) CommentAllPost(idPost int, comment *model.PostComments, pagination *model.PaginationPage) (*[]model.PostComments, error) {
	comments, err := p.usecase.CommentAllPost(idPost, comment, pagination)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (p *postHttp) PublishPost(id int) error {
	if err := p.usecase.PublishPost(id); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) RemoveComment(id int) error {
	if err := p.usecase.RemoveComment(id); err != nil {
		return err
	}
	return nil
}

func (p *postHttp) ListPublishPost(post *model.Posts, pagination *model.PaginationPage) (*[]model.Posts, error) {
	posts, err := p.usecase.ListPublishPost(post, pagination)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postHttp) ListNotPublishPost(post *model.Posts, pagination *model.PaginationPage) (*[]model.Posts, error) {
	posts, err := p.usecase.ListNotPublishPost(post, pagination)
	if err != nil {
		return nil, err
	}
	return posts, nil
}