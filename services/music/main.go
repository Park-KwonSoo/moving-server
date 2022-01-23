package music_service

import (
	"context"

	musicpb "github.com/Park-Kwonsoo/moving-server/api/protos/v1/music"

	errHandler "github.com/Park-Kwonsoo/moving-server/pkg/err-handler"

	db "github.com/Park-Kwonsoo/moving-server/models"
)

type MusicServer struct {
	musicpb.MusicServiceServer
}

/**
*	GetMusicDetail의 리턴 타입을 가져옴
 */
func getMusicDetailReturnType(e errHandler.ErrorRslt, music *db.Music) (*musicpb.GetMusicDetailRes, error) {

	if music == nil {
		return &musicpb.GetMusicDetailRes{
			RsltCd:  e.RsltCd,
			RsltMsg: e.RsltMsg,
		}, nil
	}

	return &musicpb.GetMusicDetailRes{
		RsltCd:  "00",
		RsltMsg: "Success",
		Music: &musicpb.Music{
			MusicId:     uint64(music.ID),
			TrackNumber: uint64(music.TrackNumber),

			Title:    music.Title,
			Artist:   music.Artist,
			Album:    music.Album,
			Genre:    music.Genre,
			AlbumImg: music.AlbumImg,
			MusicUrl: music.MusicUrl,

			IsTitle: music.IsTitle,
		},
	}, nil
}

/**
* GetMusicByKeyword의 리턴 타입을 가져옴
 */
func getMusicByKeywordReturnType(e errHandler.ErrorRslt, musicList []*db.Music) (*musicpb.GetMusicByKeywordRes, error) {

	if len(musicList) == 0 {
		return &musicpb.GetMusicByKeywordRes{
			RsltCd:  e.RsltCd,
			RsltMsg: e.RsltMsg,
		}, nil
	}

	rsltList := make([]*musicpb.Music, 0)
	for _, music := range musicList {
		rsltList = append(rsltList, &musicpb.Music{
			MusicId:     uint64(music.ID),
			TrackNumber: uint64(music.TrackNumber),

			Title:    music.Title,
			Artist:   music.Artist,
			Album:    music.Album,
			Genre:    music.Genre,
			AlbumImg: music.AlbumImg,
			MusicUrl: music.MusicUrl,

			IsTitle: music.IsTitle,
		})
	}

	return &musicpb.GetMusicByKeywordRes{
		RsltCd:       "00",
		RsltMsg:      "Success",
		SearchResult: rsltList,
	}, nil
}

func (s *MusicServer) GetMusicDetail(ctx context.Context, req *musicpb.GetMusicDetailReq) (*musicpb.GetMusicDetailRes, error) {

	music, err := db.FindOneMusicById(uint(req.MusicId))
	if err != nil {
		return getMusicDetailReturnType(errHandler.NotFoundErr("GetMusicDetail : Not Found Music"), nil)
	}

	return getMusicDetailReturnType(errHandler.ErrorRslt{}, music)
}

func (s *MusicServer) GetMusicByKeyword(ctx context.Context, req *musicpb.GetMusicByKeywordReq) (*musicpb.GetMusicByKeywordRes, error) {

	musicList, err := db.FindAllMusicByKeyword(req.Keyword)
	if err != nil {
		return getMusicByKeywordReturnType(errHandler.NotFoundErr("GetMusicByKeyword : Not Found Music"), nil)
	}

	return getMusicByKeywordReturnType(errHandler.ErrorRslt{}, musicList)
}