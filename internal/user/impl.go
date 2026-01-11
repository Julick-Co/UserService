package user

import (
	pb "UserService/proto/api/v1/user"
	"context"
	"fmt"
	"log"
)

type UserServiceImpl struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceImpl) SearchByNickname(ctx context.Context,
	req *pb.SearchByNicknameRequest,
) (*pb.SearchByNicknameResponse, error) {

	log.Println("Достучался до метода")

	return &pb.SearchByNicknameResponse{
		Results: []*pb.UserProfile{},
	}, nil
}

func (s *UserServiceImpl) CreateProfile(ctx context.Context,
	req *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {

	log.Println("Достучался до метода CreateProfile")

	if req.UserProfile == nil {
		return nil, fmt.Errorf("user_profile is required")
	}

	in := req.UserProfile

	return &pb.CreateProfileResponse{
		UserProfile: &pb.UserProfile{
			UserId:    "generated-id",
			Nickname:  in.Nickname,
			Bio:       in.Bio,
			AvatarUrl: in.AvatarUrl,
		},
	}, nil
}

func (s *UserServiceImpl) GetProfileByID(ctx context.Context,
	req *pb.GetProfileByIDRequest) (*pb.GetProfileByIDResponse, error) {

	log.Println("Достучался до метода GetProfileByID")

	if req.Id == "" {
		return nil, fmt.Errorf("id is required")
	}

	return &pb.GetProfileByIDResponse{
		UserProfile: &pb.UserProfile{
			UserId:    req.Id,
			Nickname:  "stored_nickname",
			Bio:       nil,
			AvatarUrl: nil,
		},
	}, nil
}

func (s *UserServiceImpl) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest,
) (*pb.UpdateProfileResponse, error) {

	log.Println("Достучался до метода")

	bio := "Bio"
	avatarUrl := "AvatarUrl"

	return &pb.UpdateProfileResponse{UserProfile: &pb.UserProfile{
		UserId:    "3213",
		Nickname:  "Nick",
		Bio:       &bio,
		AvatarUrl: &avatarUrl,
	},
	}, nil
}

func (s *UserServiceImpl) GetProfileByNickname(ctx context.Context, req *pb.GetProfileByNicknameRequest,
) (*pb.GetProfileByNicknameResponse, error) {

	log.Println("Достучался до метода")
	bio := "Bio"
	avatarUrl := "AvatarUrl"
	return &pb.GetProfileByNicknameResponse{
		UserProfile: &pb.UserProfile{
			UserId:    "123",
			Nickname:  "Nick",
			Bio:       &bio,
			AvatarUrl: &avatarUrl,
		},
	}, nil
}
