package grpcServer

import (
	"context"
	"time"
	"uid_gen/grpc/uid_gen"
	"uid_gen/internal/pkg/global"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (*GRPC_UID_server) GenUID(ctx context.Context, e *emptypb.Empty) (*uid_gen.UID, error) {
	uid := (time.Now().Unix() << (64 - 34)) | (global.MachineID << (64 - 34 - 10)) | global.GetSeq()
	return &uid_gen.UID{Val: uid}, nil
}
