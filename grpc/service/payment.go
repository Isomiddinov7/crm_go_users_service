package service

import (
	"context"
	"crm_go_users_service/config"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/grpc/client"
	"crm_go_users_service/pkg/logger"
	"crm_go_users_service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PaymentService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedPaymentServiceServer
}

func NewPaymentService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *PaymentService {
	return &PaymentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *PaymentService) Create(ctx context.Context, req *users_service.CreatePayment) (resp *users_service.Payment, err error) {

	i.log.Info("---CreatePayment------>", logger.Any("req", req))
	id, err := i.strg.Payment().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreatePayment->Payment->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		i.log.Error("!!!GetByIdTask->task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Payment().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdPayment->Payment->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (i *PaymentService) GetById(ctx context.Context, req *users_service.PaymentPrimaryKey) (resp *users_service.Payment, err error) {
	i.log.Info("---GetPaymentByID------>", logger.Any("req", req))

	resp, err = i.strg.Payment().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetPaymentByID->Payment->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *PaymentService) GetList(ctx context.Context, req *users_service.GetListPaymentRequest) (resp *users_service.GetListPaymentResponse, err error) {

	i.log.Info("---GetPayments------>", logger.Any("req", req))

	resp, err = i.strg.Payment().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetPayments->Payment->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *PaymentService) Update(ctx context.Context, req *users_service.UpdatePayment) (resp *users_service.Payment, err error) {

	i.log.Info("---UpdatePayment------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Payment().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdatePayment--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Payment().GetById(ctx, &users_service.PaymentPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *PaymentService) Delete(ctx context.Context, req *users_service.PaymentPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeletePayment------>", logger.Any("req", req))

	err = i.strg.Payment().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeletePayment->Payment->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
