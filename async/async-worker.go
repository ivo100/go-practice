package worker

//
//import (
//	"context"
//)
//
//type RequestWithContext struct {
//	CreateRequest *profile.CreateClaimRequest
//	Ctx           context.Context
//}
//
//// ChoiceRequestServiceImpl defines an implementation structure
//type ChoiceRequestServiceImpl struct {
//	Requests     chan *RequestWithContext
//	ClaimCreator common.ClaimCreator
//}
//
//// NewChoiceRequestService creates and returns an instance of ChoiceRequestServiceImpl
//// claimCreator is optional and can be null - it will default to claimservice grpc client autocreated on every call
//// unit tests can inject mock implementation
//func NewChoiceRequestService(claimCreator common.ClaimCreator) *ChoiceRequestServiceImpl {
//	zaplogger.GetDefaultLogger().Info("=== NewChoiceRequestService")
//	// channel to receive requests + context
//	requests := make(chan *RequestWithContext, 5)
//	go worker(claimCreator, requests)
//	return &ChoiceRequestServiceImpl{
//		ClaimCreator: claimCreator,
//		Requests:     requests,
//	}
//}
//
//// CreateChoiceRequest is the main entry point of the service
//// called by the user to create a claim
//// it passes via the grpc gateway and the work is scheduled asynchronously
//// more accurately name is "submit request to create user with claim"
//func (s *ChoiceRequestServiceImpl) CreateChoiceRequest(ctx context.Context, request *api.CreateChoiceRequestRequest) (*api.CreateChoiceRequestResponse, error) {
//	// minimum stateless synchronous validation of the context and request syntax
//	// fail fast on invalid request
//	rq, ct, err := common.ValidateAndInitCreateRequest(ctx, request)
//	if err != nil {
//		return nil, handleError(ctx, "Invalid request.", err)
//	}
//	// ClaimCreator is injected only by unit tests currently
//	if s.ClaimCreator == nil {
//		// preflight checks
//		if err = services.PreflightCheck(ctx, rq); err != nil {
//			return &api.CreateChoiceRequestResponse{
//				Id: request.PaymentIdentification,
//			}, err
//		}
//	}
//	// prepare request - send it via buffered channel to the worker goroutine
//	req := &RequestWithContext{
//		CreateRequest: rq,
//		Ctx:           ct,
//	}
//	s.Requests <- req
//	// here we assume that no errors will happen later in the claim service.
//	// this is not a good, just a temporary solution.
//	// we return OK regardless of the claim service result.
//	// devops should watch for BOA_RCP_CS_ERROR error string / create alert on it
//	// if it happens - it means that request may get "lost between the cracks" and
//	// it is time to think about real solution like webhook for returning real status asynchronously.
//	// another alternative is to use redis for tracking requests and return status from previous actual request.
//	return &api.CreateChoiceRequestResponse{
//		Id: request.PaymentIdentification,
//	}, nil
//}
//
//// worker reads requests from channel async execution, result is logged but ignored
//func worker(claimCreator common.ClaimCreator, requests chan *RequestWithContext) {
//	for {
//		select {
//		case req := <-requests:
//			if req != nil {
//				go asyncCreate(claimCreator, req)
//			}
//			break
//		default:
//			break
//		}
//	}
//}
//
//func asyncCreate(claimCreator common.ClaimCreator, rc *RequestWithContext) {
//	logger := zaplogger.GetLogger(rc.Ctx)
//	var err error
//	var ID string
//	ctx := rc.Ctx
//	req := rc.CreateRequest
//	if claimCreator != nil {
//		// unit test
//		ID, err = claimCreator.CreateClaim(ctx, req)
//	} else {
//		ID, err = services.CreateClaim(ctx, req)
//	}
//	if err != nil {
//		logger.Errorf("Error %v", err)
//	} else {
//		logger.Infof("Created claim %v", ID)
//	}
//}
