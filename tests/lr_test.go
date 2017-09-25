package lr_bdd_test

import (
	"encoding/json"

	"github.com/coding-yogi/go_bdd/common/utils/rest"
	"github.com/coding-yogi/go_bdd/common/utils/token"
	"github.com/coding-yogi/go_bdd/constants"
	"github.com/coding-yogi/go_bdd/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Learning Recommender Tests: ", func() {

	var accessToken string
	var url string

	Describe("Getting recommendations", func() {

		JustBeforeEach(func() {
			var err error
			url = env.BaseURL + constants.RecommenderEndpoint
			accessToken, err = token.GetOAuthToken()
			if err != nil {
				Fail("Unable to get access token")
			}
		})

		Context("For a single student", func() {
			studID := "00000034"

			reqBody := []models.RecommendationRequest{
				models.RecommendationRequest{
					StudentID: studID,
				},
			}

			It("should return all recommended courses along with topics", func() {
				reqJson, _ := json.Marshal(reqBody)
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

				recommendationRes := models.RecommenderResponse{}
				json.Unmarshal(body, &recommendationRes)

				//Expectations
				Expect(recommendationRes.RestOperationStatusVOX.Status).To(Equal("SUCCESS"))
				Expect(recommendationRes.RestOperationStatusVOX.Data.RestReturnData[0].Recommendations).NotTo(HaveLen(0))
				Expect(recommendationRes.RestOperationStatusVOX.Data.RestReturnData[0].StudID).To(Equal(studID))
			})

			Context("With Tenant Header missing", func() {
				It("should return an error", func() {
					reqJson, _ := json.Marshal(reqBody)
					headers := []rest.Header{
						rest.Header{Key: "authorization", Value: "Bearer " + accessToken},
						rest.Header{Key: "content-type", Value: "application/json"},
					}

					req := rest.GenerateRequest("POST", url, reqJson, headers)
					res, err := rest.ExecuteRequestAndGetResponse(req, client)
					if err != nil {
						Fail("Error executing recommendation request")
					}
					body := rest.GetResponseBody(res)

					recommendationRes := models.RecommenderResponse{}
					json.Unmarshal(body, &recommendationRes)

					//Expectations
					Expect(recommendationRes.RestOperationStatusVOX.Status).To(Equal("FAILURE"))
					//Expect(recommendationRes.RestOperationStatusVOX.Data.RestReturnDataAsString).To(Equal("Missing Required details to process this request"))
				})
			})

			Context("With an invalid student id", func() {
				It("should not return an error", func() {
					//Update to invalid ID
					reqBody[0].StudentID = "00000000"

					reqJson, _ := json.Marshal(reqBody)
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

					recommendationRes := models.RecommenderResponse{}
					json.Unmarshal(body, &recommendationRes)

					//Expectations
					Expect(recommendationRes.RestOperationStatusVOX.Status).To(Equal("SUCCESS"))
					Expect(recommendationRes.RestOperationStatusVOX.Data.RestReturnData[0].Recommendations).To(HaveLen(0))
					Expect(recommendationRes.RestOperationStatusVOX.Data.RestReturnData[0].StudID).To(Equal(studID))
				})
			})

		})

		Context("With invalid oauth token", func() {
			It("should return an error", func() {
				reqBody := []models.RecommendationRequest{
					models.RecommendationRequest{
						StudentID: "00000034",
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
				apiHubErrorResponse := models.APIHubErrorResponse{}
				json.Unmarshal(body, &apiHubErrorResponse)

				//Expectations
				Expect(apiHubErrorResponse.Fault.FaultString).To(Equal("Invalid Access Token"))

			})
		})
	})
})
