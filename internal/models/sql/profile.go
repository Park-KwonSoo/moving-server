package sql_model

import (
	"errors"
	"strings"

	"github.com/sirupsen/logrus"

	sqlDB "github.com/Park-Kwonsoo/moving-server/pkg/database/sql"
	getTag "github.com/Park-Kwonsoo/moving-server/pkg/get-struct-info"
	qb "github.com/Park-Kwonsoo/moving-server/pkg/query-builder"
)

type Profile struct {
	sqlDB.BaseType
	Member       Member `db:"member_mem_id varchar(255) references member(mem_id) on delete cascade" mapping:"one2one member"`
	Name         string `db:"name varchar(255)"`
	Birth        string `db:"birth varchar(10)"`
	Gender       string `db:"gender varchar(6)"`
	ProfileImage string `db:"profile_img varchar(2000)"`
}

//profile migrate
func profileMigrate() error {

	column := make([]string, 0)
	column = append(column, strings.Join(sqlDB.GetCreatedTableColumn(), ", "))
	column = append(column, getTag.GetStructInfoByTag("db", &Profile{})...)

	query := qb.CreateTable("profile").TableComlumn(
		column...,
	).ToString()

	if _, err := sqlDB.SQL.Exec(query); err != nil {
		return err
	}

	if err := sqlDB.CreateUpdateTrigger("profile"); err != nil {
		return err
	}

	err := sqlDB.TableMapping(&Profile{})
	return err
}

//새 프로필 생성
func CreateNewProfile(profile *Profile) error {

	existProfile, _ := FindOneProfileByMemberMemId(profile.Member.MemId.String)
	if existProfile != nil {
		return errors.New("Conflict")
	}

	query := qb.Insert("profile", "name, birth, gender, profile_img, member_mem_id").Value(
		profile.Name,
		profile.Birth,
		profile.Gender,
		profile.ProfileImage,
		profile.Member.MemId.String,
	).ToString()

	err := sqlDB.SQL.QueryRow(query).Scan(&profile.ID)

	return err
}

//profile id로 프로필 1개 가져오기
func FindOneProfileById(id uint) (*Profile, error) {

	profile := &Profile{}
	var memId string //유저 아이디 저장

	query := qb.Select("id, creatd_at, updated_at, name, birth, gender, profile_img, member_mem_id").From("profile").Where("id", id).ToString()
	sqlDB.SQL.QueryRow(query).Scan(
		&profile.ID, &profile.CreatedAt, &profile.UpdatedAt, &profile.Name, &profile.Birth, &profile.Gender, &profile.ProfileImage, &memId,
	)

	member, _ := FindOneMemberByMemId(memId)
	profile.Member = *member

	if !profile.Member.MemId.Valid {
		err := errors.New("Not Found")
		return nil, err
	}

	return profile, nil
}

//member id로 프로필 가져오기
func FindOneProfileByMemberMemId(memId string) (*Profile, error) {

	member, err := FindOneMemberByMemId(memId)
	if err != nil {
		return nil, err
	}
	if !member.MemId.Valid {
		err := errors.New("Not Found")
		return nil, err
	}

	profile := &Profile{}
	profile.Member = *member

	query := qb.Select("id, created_at, updated_at, name, birth, gender, profile_img").From("profile").Where("member_mem_id", memId).ToString()
	sqlDB.SQL.QueryRow(query).Scan(
		&profile.ID, &profile.CreatedAt, &profile.UpdatedAt, &profile.Name, &profile.Birth, &profile.Gender, &profile.ProfileImage,
	)

	if profile.ID == 0 {
		err := errors.New("Not Found")
		return nil, err
	}

	return profile, nil
}

func UpdateOneProfile(profile *Profile) error {

	query := qb.Update("profile").Set("name, birth, gender, profile_img", []string{
		profile.Name,
		profile.Birth,
		profile.Gender,
		profile.ProfileImage,
	}).Where("id", profile.ID).ToString()

	_, err := sqlDB.SQL.Exec(query)

	return err
}

func init() {
	if err := profileMigrate(); err != nil {
		logrus.Error(err)
	}
}
