package gostream

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lordvidex/gostream/internal/entity"

	gostreamv1 "github.com/lordvidex/gostream/pkg/api/gostream/v1"
)

// CreateUser ...
func (i *Implementation) CreateUser(ctx context.Context, req *gostreamv1.CreateUserRequest) (*gostreamv1.CreateUserResponse, error) {
	user := req.User
	id, err := i.userRepo.CreateUser(ctx, user)
	if err != nil {
		fmt.Println("error creating user", err)
		return nil, status.Errorf(codes.Internal, "error creating user: %v", err)
	}

	user.Id = id
	i.userCache.Store(entity.User{User: user})
	if err = i.publishUserUpdate(ctx, user); err != nil {
		fmt.Println("got error publishing data", err)
	}

	return &gostreamv1.CreateUserResponse{Id: id}, nil
}
