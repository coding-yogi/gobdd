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

var _ = Describe("Preference API: ", func() {

	var accessToken string
	var url string

	Describe("Setting preference", func() {

		JustBeforeEach(func() {
			var err error
			url = env.BaseURL + constants.PreferenceEndPoint
			accessToken, err = token.GetOAuthToken(env)
			if err != nil {
				Fail("Unable to get access token")
			}
		})

		Context("for combination of student and learning item", func() {

			reqBody := []requestmodels.PreferenceRequest{
				requestmodels.PreferenceRequest{
					ID:         "some_id",
					Dismissed:  false,
					Read:       true,
					Bookmarked: false,
					Liked:      false,
				},
			}
			reqJson, _ := json.Marshal(reqBody)

			It("should return success", func() {

				headers := []rest.Header{
					rest.Header{Key: "authorization", Value: "Bearer " + accessToken},
					rest.Header{Key: "content-type", Value: "application/json"},
					rest.Header{Key: "tenant_id", Value: env.Tenant},
				}

				req := rest.GenerateRequest("POST", url, reqJson, headers)
				res, err := rest.ExecuteRequestAndGetResponse(req, client)
				if err != nil {
					Fail("Error executing recommendation request")
				}
				body := rest.GetResponseBody(res)

				resObj := responsemodels.RecommenderResponse{}
				json.Unmarshal(body, &resObj)

				//Expectations
				Expect(res.StatusCode).To(Equal(200))
				Expect(resObj.RestOperationStatusVOX.Status).To(Equal("SUCCESS"))
			})

			Context("with an invalid OAuth token", func() {
				It("should return an error", func() {
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
		})

	})
})
