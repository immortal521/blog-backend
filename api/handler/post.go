package handler

import (
	"log"

	"blog-server/api/response"
	"blog-server/usecase/post"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	usecase post.UseCase
}

func RegisterPostRoute(r fiber.Router, handler *PostHandler) {
	group := r.Group("/posts")
	group.Get("/", handler.ListPosts)
}

func NewPostHandler(usecase post.UseCase) *PostHandler {
	return &PostHandler{
		usecase: usecase,
	}
}

func (h *PostHandler) ListPosts(c *fiber.Ctx) error {
	posts, err := h.usecase.ListPostsNoContent(c.Context())
	log.Println(posts)
	if err != nil {
		return err
	}
	toJ := make([]response.PostMetaRes, len(posts))
	for i, post := range posts {
		toJ[i] = response.PostMetaRes{
			ID:          post.ID,
			Title:       post.Title,
			Summary:     post.Summary,
			Cover:       post.Cover,
			ReadTime:    post.ReadTimeMinutes,
			ViewCount:   post.ViewCount,
			Status:      string(post.Status),
			PublishedAt: post.PublishedAt,
			UpdatedAt:   post.UpdatedAt,
			Author:      post.Edges.Author.Username,

			Categories: func() []*response.PostCategoryRes {
				cats := make([]*response.PostCategoryRes, len(post.Edges.Categories))
				for i, c := range post.Edges.Categories {
					cats[i] = &response.PostCategoryRes{
						ID:   c.ID,
						Name: c.Name,
					}
				}
				return cats
			}(),
			Tags: func() []*response.PostTagRes {
				tags := make([]*response.PostTagRes, len(post.Edges.Tags))
				for i, t := range post.Edges.Tags {
					tags[i] = &response.PostTagRes{
						ID:   t.ID,
						Name: t.Name,
					}
				}
				return tags
			}(),
		}
	}
	result := response.Success(toJ)
	return c.JSON(result)
}
