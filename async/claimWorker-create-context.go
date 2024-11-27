package worker

//
//import (
//	"context"
//	"sync"
//	"time"
//
//	"github.com/spf13/viper"
//)
//
//var logger = zaplogger.GetDefaultLogger()
//
//type DataRepository interface {
//	UpdatePaymentRequest(pr *db.PaymentRequest) error
//	GetOutstandingPaymentRequests() ([]*db.PaymentRequest, error)
//}
//
//type ClaimWorker struct {
//	dataRepo      DataRepository
//	contextGetter authnz.ContextGetter
//}
//
//func NewClaimWorker(dataRepo DataRepository) *ClaimWorker {
//	return &ClaimWorker{
//		dataRepo:      dataRepo,
//		contextGetter: authnz.DefaultContextGetter(),
//	}
//}
//
//func (cw *ClaimWorker) CheckAndCreateClaims(ctx context.Context, wg *sync.WaitGroup) {
//	defer wg.Done()
//	timer := time.NewTicker(time.Duration(viper.GetInt("claimPollInterval")) * time.Second)
//	for {
//		select {
//		case <-timer.C:
//			cw.checkAndCreateUnprocessedClaims()
//		case <-ctx.Done():
//			return
//		}
//	}
//}
//
//func (cw *ClaimWorker) checkAndCreateUnprocessedClaims() {
//	unprocessedClaims, err := cw.dataRepo.GetOutstandingPaymentRequests()
//	if err != nil {
//		logger.Errorf("unable to retrieve unprocessed payment requests from database : %v", err)
//	}
//
//	if len(unprocessedClaims) == 0 {
//		logger.Debugf("no unprocessed records. Next check in %d seconds", viper.GetInt("claimPollInterval"))
//	}
//
//	for _, unprocessedClaim := range unprocessedClaims {
//		originalRequest, err := db.ChoiceRequestBlob(unprocessedClaim.Content).ToAPIModel()
//		if err != nil {
//			logger.Errorf("unable to unmarshall request with ID '%v' : %v", unprocessedClaim.ID, err)
//			continue
//		}
//		outboundClaimRequest, err := common.ConvertToCreateClaimRequest(unprocessedClaim.ClientID, unprocessedClaim.Realm, originalRequest)
//		if err != nil {
//			logger.Errorf("unable to convert claim request with ID '%v' : %v", unprocessedClaim.ID, err)
//			continue
//		}
//		outCtx, err := cw.contextGetter.GetOutgoingSystemContext(authnz.CreateContextByRealm(unprocessedClaim.Realm))
//		if err != nil {
//			logger.Errorf("unable to get outgoing system context : %v", err)
//			continue
//		}
//		profileClientWrapper, err := profileclient.GetClientWrapper(outCtx)
//		if err != nil {
//			logger.Errorf("unable to get Client Wrapper : %v", err)
//			continue
//		}
//		resp, err := profileClientWrapper.ClaimServiceClient.CreateClaim(outCtx, outboundClaimRequest)
//		if err != nil {
//			logger.Errorf("unable to process claim with internal ID '%v' and Reference Number '%v' : %v", unprocessedClaim.ID, outboundClaimRequest.GetClaim().GetReferenceNumber(), err)
//			unprocessedClaim.Attempts++
//		} else {
//			unprocessedClaim.Processed = true
//			logger.Infof("Created claim with ID '%v' from queue entry with ID '%v'", resp.GetData().GetId(), unprocessedClaim.ID)
//		}
//		err = cw.dataRepo.UpdatePaymentRequest(unprocessedClaim)
//		if err != nil {
//			logger.Errorf("unable to mark claim with ID '%v' as processed : %v", unprocessedClaim.ID, err)
//			continue
//		}
//	}
//}
