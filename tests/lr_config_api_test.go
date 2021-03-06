package lr_bdd_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.wdf.sap.corp/ml-base/lr-bdd-tests/common/utils/rest"
	"github.wdf.sap.corp/ml-base/lr-bdd-tests/common/utils/token"
	"github.wdf.sap.corp/ml-base/lr-bdd-tests/constants"
	"github.wdf.sap.corp/ml-base/lr-bdd-tests/models/api/request"
	"github.wdf.sap.corp/ml-base/lr-bdd-tests/models/api/response"
)

var _ = Describe("Config API: ", func() {

	var accessToken string
	var url string

	Describe("Updating configuration", func() {

		JustBeforeEach(func() {
			var err error
			url = env.BaseURL + constants.ConfigEndPoint
			accessToken, err = token.GetOAuthToken(env)
			if err != nil {
				Fail("Unable to get access token")
			}
		})

		Context("to (isEnabled) true for a tenant", func() {
			It("should return proper response", func() {
				reqBody := requestmodels.ConfigRequest{
					Main: requestmodels.ConfigRequestMain{
						IsEnabled: true,
					},
				}
				reqJson, _ := json.Marshal(reqBody)
				headers := []rest.Header{
					rest.Header{Key: "authorization", Value: "Bearer " + accessToken},
					rest.Header{Key: "content-type", Value: "application/json"},
					rest.Header{Key: "tenant_id", Value: env.Tenant},
				}

				req := rest.GenerateRequest("PUT", url, reqJson, headers)
				res, err := rest.ExecuteRequestAndGetResponse(req, client)
				if err != nil {
					Fail("Error executing recommendation request")
				}
				body := rest.GetResponseBody(res)

				resObj := responsemodels.RecommenderResponse{}
				json.Unmarshal(body, &resObj)

				var retData string
				json.Unmarshal(resObj.RestOperationStatusVOX.Data.RestReturnData, &retData)

				//Expectations
				Expect(retData).To(Equal("Available"))
				Expect(resObj.RestOperationStatusVOX.Status).To(Equal("SUCCESS"))
			})

		})

		Context("with an invalid OAuth token", func() {
			It("should return an error", func() {
				reqBody := requestmodels.ConfigRequest{
					Main: requestmodels.ConfigRequestMain{
						IsEnabled: true,
					},
				}
				reqJson, _ := json.Marshal(reqBody)

				someInvalidAccessToken := "a1ABCDEFbcdGhijkLMnOp6QR01Ts"
				headers := []rest.Header{
					rest.Header{Key: "authorization", Value: "Bearer " + someInvalidAccessToken},
					rest.Header{Key: "content-type", Value: "application/json"},
					rest.Header{Key: "tenant_id", Value: env.Tenant},
				}

				req := rest.GenerateRequest("POST", url, reqJson, headers)
				res, err := rest.ExecuteRequestAndGetResponse(req, client)
				if err != nil {
					Fail("Error executing recommendation request")
				}

				body := rest.GetResponseBody(res)
				resObj := responsemodels.APIHubErrorResponse{}
				json.Unmarshal(body, &resObj)

				//Expectations
				Expect(res.StatusCode).To(Equal(401))
				Expect(resObj.Fault.FaultString).To(Equal("Invalid Access Token"))
			})
		})

		Context("for an invalid Tenant ID", func() {
			It("should return an error", func() {
				reqBody := requestmodels.ConfigRequest{
					Main: requestmodels.ConfigRequestMain{
						IsEnabled: true,
					},
				}
				reqJson, _ := json.Marshal(reqBody)

				headers := []rest.Header{
					rest.Header{Key: "authorization", Value: "Bearer " + accessToken},
					rest.Header{Key: "content-type", Value: "application/json"},
					rest.Header{Key: "tenant_id", Value: "some_non_existant_tenant"},
				}

				req := rest.GenerateRequest("PUT", url, reqJson, headers)
				res, err := rest.ExecuteRequestAndGetResponse(req, client)
				if err != nil {
					Fail("Error executing recommendation request")
				}

				body := rest.GetResponseBody(res)
				resObj := responsemodels.RecommenderResponse{}
				json.Unmarshal(body, &resObj)

				var errorString string
				json.Unmarshal(resObj.RestOperationStatusVOX.Data.RestReturnData, &errorString)

				//Expectations
				Expect(res.StatusCode).To(Equal(400))
				Expect(errorString).To(Equal("Invalid Tenant"))
			})
		})
	})
})
