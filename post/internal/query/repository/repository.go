package repository

import (
	"errors"
	"github.com/JieeiroSst/itjob/model"
	"gorm.io/gorm"
	"time"
)

type postsRepository struct {
	db *gorm.DB
}

type PostsRepository interface {
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

func NewPostsRepository(db *gorm.DB) PostsRepository {
	return &postsRepository{
		db:db,
	}
}

func (p *postsRepository) CreatePosts(posts *model.Posts) error {
	if err := p.db.Create(&posts).Error; err != nil{
		return errors.New("create failed")
	}
	return nil
}

func (p *postsRepository) UpdatePosts(id int, posts *model.Posts) error {
	if err := p.db.Model(model.Posts{}).Where("id = ? ", id).Updates(posts).Error; err != nil{
		return errors.New("update failed")
	}
	return nil
}

func (p *postsRepository) DeletePosts(id int) error {
	if err := p.db.Delete(model.Posts{}, "id = ?", id).Error; err != nil{
		return errors.New("delete failed")
	}
	return nil
}

func (p *postsRepository) PostsById(id int) (*model.Posts, error) {
	var posts model.Posts
	result := p.db.Where("id = ?", id).Find(&posts)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &posts, nil
}

func (p *postsRepository) CreateProfile(profile *model.Profiles) error {
	if err := p.db.Create(&profile).Error; err != nil{
		return errors.New("create failed")
	}
	return nil
}

func (p *postsRepository) UpdateProfile(id int, profile *model.Profiles) error {
	if err := p.db.Model(model.Profiles{}).Where("id = ? ", id).Updates(profile).Error; err != nil{
		return errors.New("update failed")
	}
	return nil
}

func (p *postsRepository) ProfileById(id int) (*model.Profiles, error) {
	var profile model.Profiles
	result := p.db.Where("id = ?", id).Find(&profile)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &profile, nil
}

func (p *postsRepository) CreatePostMetas(metas *model.PostMetas) error {
	if err := p.db.Create(&metas).Error; err != nil{
		return errors.New("create failed")
	}
	return nil
}

func (p *postsRepository) UpdatePostMetas(id int, metas *model.PostMetas) error {
	if err := p.db.Model(model.PostMetas{}).Where("id = ? ", id).Updates(metas).Error; err != nil{
		return errors.New("update failed")
	}
	return nil
}

func (p *postsRepository) DeletePostMetas(id int) error {
	if err := p.db.Delete(model.PostMetas{}, "id = ?", id).Error; err != nil{
		return errors.New("delete failed")
	}
	return nil
}

func (p *postsRepository) PostMetasById(id int) (*model.PostMetas, error) {
	var postMeta model.PostMetas
	result := p.db.Where("id = ?", id).Find(&postMeta)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &postMeta, nil
}

func (p *postsRepository) CreateComment(comment *model.PostComments) error {
	if err := p.db.Create(&comment).Error; err != nil{
		return errors.New("create failed")
	}
	return nil
}

func (p *postsRepository) CreateCategories(categories *model.Categories) error {
	if err := p.db.Create(&categories).Error; err != nil{
		return errors.New("create failed")
	}
	return nil
}

func (p *postsRepository) UpdateCategories(id int, categories *model.Categories) error {
	if err := p.db.Model(model.Categories{}).Where("id = ? ", id).Updates(categories).Error; err != nil{
		return errors.New("update failed")
	}
	return nil
}

func (p *postsRepository) DeleteCategories(id int) error {
	if err := p.db.Delete(model.Categories{}, "id = ?", id).Error; err != nil{
		return errors.New("delete failed")
	}
	return nil
}

func (p *postsRepository) CategoriesById(id int) (*model.Categories, error) {
	var category model.Categories
	result := p.db.Where("id = ?", id).Find(&category)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &category, nil
}

func (p *postsRepository) PostsAll(post *model.Posts, pagination *model.PaginationPage) (*[]model.Posts, error) {
	var posts []model.Posts
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := p.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&model.Posts{}).Preload("PostMetas").Preload("PostComments").Preload("Categories").Where(post).Find(&posts)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &posts, nil
}

func (p *postsRepository) ProfileAll(profile *model.Profiles, pagination *model.PaginationPage) (*[]model.Profiles, error) {
	var profiles []model.Profiles
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := p.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&model.Posts{}).Preload("PostComments").Where(profile).Find(&profiles)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &profiles, nil
}

func (p *postsRepository) PostMetasAll(postMeta *model.PostMetas, pagination *model.PaginationPage) (*[]model.PostMetas, error) {
	var postMetas []model.PostMetas
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := p.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&model.Posts{}).Where(postMeta).Find(&postMetas)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &postMetas, nil
}

func (p *postsRepository) CategoriesAll(category *model.Categories, pagination *model.PaginationPage) (*[]model.Categories, error) {
	var categories []model.Categories
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := p.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&model.Posts{}).Where(category).Find(&categories)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &categories, nil
}

func (p *postsRepository) CommentAllPost(idPost int,comment *model.PostComments, pagination *model.PaginationPage) (*[]model.PostComments, error) {
	var comments []model.PostComments
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := p.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&model.Posts{}).Where(comment).Where("postId =? ",idPost).Find(&comments)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &comments, nil
}

func (p *postsRepository) PublishPost(id int) error {
	if err := p.db.Model(model.Posts{}).Where("id = ? ", id).UpdateColumns(model.Posts{Published:1,PublishedAt:time.Now()}).Error; err != nil{
		return errors.New("update failed")
	}
	return nil
}

func (p *postsRepository) RemoveComment(id int) error {
	if err := p.db.Model(model.Posts{}).Where("id = ? ", id).Update("published", 0).Error; err != nil{
		return errors.New("remove failed")
	}
	return nil
}

func (p *postsRepository) ListPublishPost(post *model.Posts, pagination *model.PaginationPage) (*[]model.Posts, error) {
	var posts []model.Posts
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := p.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&model.Posts{}).Preload("PostMetas").Preload("PostComments").Preload("Categories").
																		Where(post).Where("published = ?" ,1).Find(&posts)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &posts, nil
}

func (p *postsRepository) ListNotPublishPost(post *model.Posts, pagination *model.PaginationPage) (*[]model.Posts, error) {
	var posts []model.Posts
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := p.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&model.Posts{}).Preload("PostMetas").Preload("PostComments").Preload("Categories").
																	Where(post).Where("published = ?" ,0).Find(&posts)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &posts, nil
}
