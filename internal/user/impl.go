package user

import (
	pb "UserService/proto/pkg/api/user"
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
	req *pb.CreateProfileRequest) (*pb.UserProfile, error) {

	log.Println("Достучался до метода CreateProfile")

	if req.UserProfile == nil {
		return nil, fmt.Errorf("user_profile is required")
	}

	in := req.UserProfile

	return &pb.UserProfile{
		UserId:    "generated-id",
		Nickname:  in.Nickname,
		Bio:       in.Bio,
		AvatarUrl: in.AvatarUrl,
	}, nil
}

func (s *UserServiceImpl) GetProfileByID(ctx context.Context,
	req *pb.GetProfileByIDRequest) (*pb.UserProfile, error) {

	log.Println("Достучался до метода GetProfileByID")

	if req.Id == "" {
		return nil, fmt.Errorf("id is required")
	}

	return &pb.UserProfile{
		UserId:    req.Id,
		Nickname:  "stored_nickname",
		Bio:       nil,
		AvatarUrl: nil,
	}, nil
}
