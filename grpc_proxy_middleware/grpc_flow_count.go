package grpc_proxy_middleware

import (
	"go_gateway/dao"
	"go_gateway/public"
	"log"

	"google.golang.org/grpc"
)

func GrpcFlowCountMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		totalCounter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
		if err != nil {
			return err
		}
		totalCounter.Increase()
		serviceCounter, err := public.FlowCounterHandler.GetCounter(public.FlowServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			return err
		}
		serviceCounter.Increase()

		if err := handler(srv, ss); err != nil {
			log.Printf("GrpcFlowCountMiddleware failed with error %v\n", err)
			return err
		}
		return nil
	}
}
