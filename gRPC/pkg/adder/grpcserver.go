package adder

import "context"

//GRPCServer
type GRPCServer struct {
}

//Add
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.Getx() + req.Gety()}, nil
}
