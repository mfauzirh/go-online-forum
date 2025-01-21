package posts

import (
	"context"

	"github.com/mfauzirh/go-online-forum/internal/model/posts"
)

func (r *repository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := `INSERT INTO comments (post_id, user_id, comment_content, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.CommentContent, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetCommentByPostId(ctx context.Context, postID int) ([]posts.Comment, error) {
	query := `SELECT c.id, c.user_id, c.comment_content, u.username FROM comments c JOIN users u ON c.user_id = u.id WHERE c.post_id = ?`

	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	response := make([]posts.Comment, 0)
	for rows.Next() {
		var comment posts.Comment

		err := rows.Scan(&comment.ID, &comment.UserID, &comment.CommentContent, &comment.Username)
		if err != nil {
			return nil, err
		}

		response = append(response, comment)
	}

	return response, nil
}
