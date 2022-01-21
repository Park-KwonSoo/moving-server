package member_service

import (
	"context"

	memberpb "github.com/Park-Kwonsoo/moving-server/api/protos/v1/member"

	errHandler "github.com/Park-Kwonsoo/moving-server/pkg/err-handler"

	db "github.com/Park-Kwonsoo/moving-server/models"
	jwtUtil "github.com/Park-Kwonsoo/moving-server/pkg/jwt-utility"
)

type MemberServer struct {
	memberpb.MemberServiceServer
}

//profile 리턴 값을 가져오는 메서드
func getProfileReturnType(e errHandler.ErrorRslt, profile *db.Profile) (*memberpb.GetMyProfileRes, error) {

	if profile == nil {
		return &memberpb.GetMyProfileRes{
			RsltCd:    e.RsltCd,
			RsltMsg:   e.RsltMsg,
			MyProfile: nil,
		}, nil
	}

	return &memberpb.GetMyProfileRes{
		RsltCd:  "00",
		RsltMsg: "Success",
		MyProfile: &memberpb.Profile{
			Id:        uint64(profile.ID),
			CreatedAt: profile.CreatedAt.String(),
			UpdatedAt: profile.UpdatedAt.String(),
			DeletedAt: profile.DeletedAt.String(),

			Member: &memberpb.Member{
				Id:        uint64(profile.Member.ID),
				CreatedAt: profile.Member.CreatedAt.String(),
				UpdatedAt: profile.Member.UpdatedAt.String(),
				DeletedAt: profile.Member.DeletedAt.String(),

				MemId:   profile.Member.MemId.String,
				MemType: profile.Member.MemType,
			},

			Name:       profile.Name,
			Gender:     profile.Gender,
			Birth:      profile.Birth,
			ProfileImg: profile.ProfileImage,
		},
	}, nil
}

/*
* Get My Profile
 */
func (s *MemberServer) GetMyProfile(ctx context.Context, req *memberpb.GetMyProfileReq) (*memberpb.GetMyProfileRes, error) {
	memId, err := jwtUtil.ValidateToken(req.Token)
	if err != nil {
		return getProfileReturnType(errHandler.AuthorizedErr(), nil)
	}

	profile, err := db.FindProfileByMemberMemId(memId)
	if profile == nil {
		return getProfileReturnType(errHandler.NotFoundErr(), nil)
	}

	return getProfileReturnType(errHandler.ErrorRslt{}, profile)
}

/*
* update my profile
 */
func (s *MemberServer) UpdateMyProfile(ctx context.Context, req *memberpb.UpdateMyProfileReq) (*memberpb.UpdateMyProfileRes, error) {

	memId, err := jwtUtil.ValidateToken(req.Token)
	if err != nil {
		return &memberpb.UpdateMyProfileRes{
			RsltCd:  "01",
			RsltMsg: "Authorization",
		}, nil
	}

	profile, err := db.FindProfileByMemberMemId(memId)
	if profile == nil {
		return &memberpb.UpdateMyProfileRes{
			RsltCd:  "04",
			RsltMsg: "NotFound",
		}, nil
	}

	//profile 이름 업데이트
	if len(req.Name) > 0 {
		profile.Name = req.Name
	}

	//profile 생일 업데이트
	if len(req.Birth) > 0 {
		profile.Birth = req.Birth
	}

	//profile 성별 업데이트
	if len(req.Gender) > 0 {
		profile.Gender = req.Gender
	}

	//profile 사진 업데이트
	if len(req.ProfileImg) > 0 {
		profile.ProfileImage = req.ProfileImg
	}

	err = db.UpdateOneProfile(profile)
	if err != nil {
		return &memberpb.UpdateMyProfileRes{
			RsltCd:  "99",
			RsltMsg: "Failure",
		}, nil
	}

	return &memberpb.UpdateMyProfileRes{
		RsltCd:  "00",
		RsltMsg: "Success",
	}, nil
}

/*
* Get My Playlist
 */

func (s *MemberServer) GetMyPlaylist(ctx context.Context, req *memberpb.GetMyPlaylistReq) (*memberpb.GetMyPlaylistRes, error) {
	return &memberpb.GetMyPlaylistRes{
		RsltCd:  "00",
		RsltMsg: "Success",
	}, nil
}
