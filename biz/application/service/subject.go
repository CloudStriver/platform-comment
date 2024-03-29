package service

import (
	"context"
	"fmt"
	"github.com/CloudStriver/go-pkg/utils/util/log"
	"github.com/CloudStriver/platform-comment/biz/infrastructure/consts"
	"github.com/CloudStriver/platform-comment/biz/infrastructure/convertor"
	subjectMapper "github.com/CloudStriver/platform-comment/biz/infrastructure/mapper/subject"
	gencomment "github.com/CloudStriver/service-idl-gen-go/kitex_gen/platform/comment"
	"github.com/google/wire"
)

type ISubjectService interface {
	UpdateCount(ctx context.Context, rootId, subjectId, fatherId string, count int64)
	GetCommentSubject(ctx context.Context, req *gencomment.GetCommentSubjectReq) (resp *gencomment.GetCommentSubjectResp, err error)
	CreateCommentSubject(ctx context.Context, req *gencomment.CreateCommentSubjectReq) (resp *gencomment.CreateCommentSubjectResp, err error)
	UpdateCommentSubject(ctx context.Context, req *gencomment.UpdateCommentSubjectReq) (resp *gencomment.UpdateCommentSubjectResp, err error)
	DeleteCommentSubject(ctx context.Context, req *gencomment.DeleteCommentSubjectReq) (resp *gencomment.DeleteCommentSubjectResp, err error)
}

type SubjectService struct {
	SubjectMongoMapper subjectMapper.IMongoMapper
}

var SubjectSet = wire.NewSet(
	wire.Struct(new(SubjectService), "*"),
	wire.Bind(new(ISubjectService), new(*SubjectService)),
)

func (s *SubjectService) GetCommentSubject(ctx context.Context, req *gencomment.GetCommentSubjectReq) (resp *gencomment.GetCommentSubjectResp, err error) {
	resp = new(gencomment.GetCommentSubjectResp)
	var data *subjectMapper.Subject
	fmt.Printf("[%v]\n", req.Id)
	if data, err = s.SubjectMongoMapper.FindOne(ctx, req.Id); err != nil {
		log.CtxError(ctx, "获取评论区详情 失败[%v]\n", err)
		return resp, err
	}
	resp.Subject = convertor.SubjectMapperToSubjectDetail(data)
	return resp, nil
}

func (s *SubjectService) CreateCommentSubject(ctx context.Context, req *gencomment.CreateCommentSubjectReq) (resp *gencomment.CreateCommentSubjectResp, err error) {
	resp = new(gencomment.CreateCommentSubjectResp)
	data := convertor.SubjectToSubjectMapper(req.Subject)
	if resp.Id, err = s.SubjectMongoMapper.Insert(ctx, data); err != nil {
		log.CtxError(ctx, "创建评论区 失败[%v]\n", err)
		return resp, err
	}
	return resp, nil
}

func (s *SubjectService) UpdateCount(ctx context.Context, rootId, subjectId, fatherId string, count int64) {
	if rootId == subjectId {
		// 一级评论
		if fatherId == subjectId {
			s.SubjectMongoMapper.UpdateCount(ctx, subjectId, count, count)
		}
	} else {
		// 二级评论 + 三级评论
		if fatherId != subjectId {
			s.SubjectMongoMapper.UpdateCount(ctx, subjectId, count, consts.InitNumber)
		}
	}
}

func (s *SubjectService) UpdateCommentSubject(ctx context.Context, req *gencomment.UpdateCommentSubjectReq) (resp *gencomment.UpdateCommentSubjectResp, err error) {
	resp = new(gencomment.UpdateCommentSubjectResp)
	data := convertor.SubjectToSubjectMapper(req.Subject)
	if _, err = s.SubjectMongoMapper.Update(ctx, data); err != nil {
		log.CtxError(ctx, "修改评论区信息 失败[%v]\n", err)
		return resp, err
	}
	return resp, nil
}

func (s *SubjectService) DeleteCommentSubject(ctx context.Context, req *gencomment.DeleteCommentSubjectReq) (resp *gencomment.DeleteCommentSubjectResp, err error) {
	resp = new(gencomment.DeleteCommentSubjectResp)
	if _, err = s.SubjectMongoMapper.Delete(ctx, req.Id); err != nil {
		log.CtxError(ctx, "删除评论区 失败[%v]\n", err)
		return resp, err
	}
	return resp, nil
}
