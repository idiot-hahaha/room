package grpc

//type GrpcServer struct {
//	service *service.DanmakuReplyService
//	api2.UnimplementedReplyServerServer
//}
//
//func NewGrpcServer(service *service.DanmakuReplyService, conf *model.GrpcConfig) (server *GrpcServer, err error) {
//	lis, err := net.Listen("tcp", conf.Host)
//	if err != nil {
//		return nil, err
//	}
//	server = &GrpcServer{
//		service: service,
//	}
//	grpcServer := grpc.NewServer()
//	grpc := &GrpcServer{
//		service: service,
//	}
//	api2.RegisterReplyServerServer(grpcServer, grpc)
//	go func() {
//		if err = grpcServer.Serve(lis); err != nil {
//			panic(err)
//		}
//	}()
//	return
//}
//
//func (s *GrpcServer) Close() error {
//	return nil
//}
//
//func (s *GrpcServer) Ping(ctx context.Context, req *api2.PingReq) (res *api2.PingRes, err error) {
//	return s.service.Ping(ctx, req)
//}
//
//func (s *GrpcServer) ReReplyByGroupID(ctx context.Context, req *api2.ReplyByGroupIDReq) (res *api2.ReplyByGroupIDRes, err error) {
//	return s.service.ReplyByGroupID(ctx, req)
//}
