package adaptor

import (
	"context"
	"github.com/CloudStriver/platform-comment/biz/application/service"
	"github.com/CloudStriver/platform-comment/biz/infrastructure/config"
	"github.com/CloudStriver/service-idl-gen-go/kitex_gen/platform/comment"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/mr"
)

type CommentServerImpl struct {
	*config.Config
	CommentService service.ICommentService
	SubjectService service.ISubjectService
}

func (c *CommentServerImpl) GetComment(ctx context.Context, req *comment.GetCommentReq) (res *comment.GetCommentResp, err error) {
	return c.CommentService.GetComment(ctx, req)
}

func (c *CommentServerImpl) GetCommentList(ctx context.Context, req *comment.GetCommentListReq) (res *comment.GetCommentListResp, err error) {
	return c.CommentService.GetCommentList(ctx, req)
}

func (c *CommentServerImpl) CreateComment(ctx context.Context, req *comment.CreateCommentReq) (res *comment.CreateCommentResp, err error) {
	if res, err = c.CommentService.CreateComment(ctx, req); err != nil {
		return res, err
	}
	_ = mr.Finish(func() error {
		c.CommentService.UpdateAfterCreateComment(ctx, req.Comment)
		return nil
	}, func() error {
		c.SubjectService.UpdateAfterCreateComment(ctx, req.Comment)
		return nil
	})
	return res, nil
}

func (c *CommentServerImpl) CreateAfter(data *comment.Comment) {

}

func (c *CommentServerImpl) UpdateComment(ctx context.Context, req *comment.UpdateCommentReq) (res *comment.UpdateCommentResp, err error) {
	return c.CommentService.UpdateComment(ctx, req)
}

func (c *CommentServerImpl) DeleteComment(ctx context.Context, req *comment.DeleteCommentReq) (res *comment.DeleteCommentResp, err error) {
	return c.CommentService.DeleteComment(ctx, req)
}

func (c *CommentServerImpl) DeleteCommentWithUserId(ctx context.Context, req *comment.DeleteCommentWithUserIdReq) (res *comment.DeleteCommentWithUserIdResp, err error) {
	return c.CommentService.DeleteCommentWithUserId(ctx, req)
}

func (c *CommentServerImpl) SetCommentState(ctx context.Context, req *comment.SetCommentStateReq) (res *comment.SetCommentStateResp, err error) {
	return c.CommentService.SetCommentState(ctx, req)
}

func (c *CommentServerImpl) SetCommentAttrs(ctx context.Context, req *comment.SetCommentAttrsReq) (res *comment.SetCommentAttrsResp, err error) {
	var resp *comment.GetCommentSubjectResp
	if resp, err = c.SubjectService.GetCommentSubject(ctx, &comment.GetCommentSubjectReq{FilterOptions: &comment.SubjectFilterOptions{OnlySubjectId: lo.ToPtr(req.SubjectId), OnlyUserId: lo.ToPtr(req.UserId)}}); err != nil {
		return res, err
	}
	return c.CommentService.SetCommentAttrs(ctx, req, resp)
}

func (c *CommentServerImpl) GetCommentSubject(ctx context.Context, req *comment.GetCommentSubjectReq) (res *comment.GetCommentSubjectResp, err error) {
	return c.SubjectService.GetCommentSubject(ctx, req)
}

func (c *CommentServerImpl) CreateCommentSubject(ctx context.Context, req *comment.CreateCommentSubjectReq) (res *comment.CreateCommentSubjectResp, err error) {
	return c.SubjectService.CreateCommentSubject(ctx, req)
}

func (c *CommentServerImpl) UpdateCommentSubject(ctx context.Context, req *comment.UpdateCommentSubjectReq) (res *comment.UpdateCommentSubjectResp, err error) {
	return c.SubjectService.UpdateCommentSubject(ctx, req)
}

func (c *CommentServerImpl) DeleteCommentSubject(ctx context.Context, req *comment.DeleteCommentSubjectReq) (res *comment.DeleteCommentSubjectResp, err error) {
	return c.SubjectService.DeleteCommentSubject(ctx, req)
}

func (c *CommentServerImpl) SetCommentSubjectState(ctx context.Context, req *comment.SetCommentSubjectStateReq) (res *comment.SetCommentSubjectStateResp, err error) {
	return c.SubjectService.SetCommentSubjectState(ctx, req)
}

func (c *CommentServerImpl) SetCommentSubjectAttrs(ctx context.Context, req *comment.SetCommentSubjectAttrsReq) (res *comment.SetCommentSubjectAttrsResp, err error) {
	return c.SubjectService.SetCommentSubjectAttrs(ctx, req)
}
