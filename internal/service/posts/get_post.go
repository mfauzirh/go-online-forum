package posts

import (
	"context"

	"github.com/mfauzirh/go-online-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id from database")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikeByPostId(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error count like in database")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentByPostId(ctx, int(postID))
	if err != nil {
		log.Error().Err(err).Msg("error get comment for post from database")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:           postDetail.ID,
			UserID:       postDetail.UserID,
			Username:     postDetail.Username,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil
}
