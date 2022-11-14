package grpcServer

import "uid_gen/grpc/uid_gen"

type GRPC_UID_server struct {
	uid_gen.UnimplementedUidGenServer
}
