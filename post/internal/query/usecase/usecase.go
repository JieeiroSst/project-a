package usecase

import (
	"github.com/JieeiroSst/itjob/model"
	"github.com/JieeiroSst/itjob/post/internal/query/repository"
)

type postUsecase struct {
	repository repository.PostsRepository
}

type PostUsecase interface {
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

func NewPostUsecase(repository repository.PostsRepository) PostUsecase {
	return &postUsecase{
		repository:repository,
	}
}

func (p *postUsecase) CreatePosts(posts *model.Posts) error {
	if err := p.repository.CreatePosts(posts); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) UpdatePosts(id int, posts *model.Posts) error {
	if err := p.repository.UpdatePosts(id, posts); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) DeletePosts(id int) error {
	if err := p.repository.DeletePosts(id); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) PostsById(id int) (*model.Posts, error) {
	post, err := p.repository.PostsById(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *postUsecase) CreateProfile(profile *model.Profiles) error {
	if err := p.repository.CreateProfile(profile); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) UpdateProfile(id int, profile *model.Profiles) error {
	if err := p.repository.UpdateProfile(id, profile); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) ProfileById(id int) (*model.Profiles, error) {
	profile, err := p.repository.ProfileById(id)
	if err != nil {
		return nil ,err
	}
	return profile, nil 
}

func (p *postUsecase) CreatePostMetas(metas *model.PostMetas) error {
	if err := p.repository.CreatePostMetas(metas); err != nil {
		return err 
	}
	return nil
}

func (p *postUsecase) UpdatePostMetas(id int, metas *model.PostMetas) error {
	if err := p.repository.UpdatePostMetas(id, metas); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) DeletePostMetas(id int) error {
	if err := p.repository.DeletePostMetas(id); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) PostMetasById(id int) (*model.PostMetas, error) {
	postMetas, err := p.repository.PostMetasById(id)
	if err != nil {
		return nil ,err
	}
	return postMetas, nil
}

func (p *postUsecase) CreateComment(comment *model.PostComments) error {
	if err := p.repository.CreateComment(comment); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) CreateCategories(categories *model.Categories) error {
	if err := p.repository.CreateCategories(categories); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) UpdateCategories(id int, categories *model.Categories) error {
	if err := p.repository.UpdateCategories(id, categories); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) DeleteCategories(id int) error {
	if err := p.repository.DeleteCategories(id); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) CategoriesById(id int) (*model.Categories, error) {
	category, err := p.repository.CategoriesById(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (p *postUsecase) PostsAll(post *model.Posts, pagination *model.PaginationPage) (*[]model.Posts, error) {
	posts, err := p.repository.PostsAll(post, pagination)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postUsecase) ProfileAll(profile *model.Profiles, pagination *model.PaginationPage) (*[]model.Profiles, error) {
	profiles, err := p.repository.ProfileAll(profile, pagination)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (p *postUsecase) PostMetasAll(postMetas *model.PostMetas, pagination *model.PaginationPage) (*[]model.PostMetas, error) {
	postMeta, err := p.repository.PostMetasAll(postMetas, pagination)
	if err != nil {
		return nil ,err
	}
	return postMeta, nil
}

func (p *postUsecase) CategoriesAll(category *model.Categories, pagination *model.PaginationPage) (*[]model.Categories, error) {
	categories, err := p.repository.CategoriesAll(category, pagination)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (p *postUsecase) CommentAllPost(idPost int, comment *model.PostComments, pagination *model.PaginationPage) (*[]model.PostComments, error) {
	comments, err := p.repository.CommentAllPost(idPost, comment, pagination)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (p *postUsecase) PublishPost(id int) error {
	if err := p.repository.PublishPost(id); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) RemoveComment(id int) error {
	if err := p.repository.RemoveComment(id); err != nil {
		return err
	}
	return nil
}

func (p *postUsecase) ListPublishPost(post *model.Posts, pagination *model.PaginationPage) (*[]model.Posts, error) {
	posts, err := p.repository.ListPublishPost(post, pagination)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postUsecase) ListNotPublishPost(post *model.Posts, pagination *model.PaginationPage) (*[]model.Posts, error) {
	posts, err := p.repository.ListNotPublishPost(post, pagination)
	if err != nil {
		return nil, err
	}
	return posts, nil
}