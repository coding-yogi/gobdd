package lr_bdd_test

import (
	"encoding/json"

	"github.com/leonelquinteros/gorand"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.wdf.sap.corp/ml-base/lr-bdd-tests/common/utils/rest"
	"github.wdf.sap.corp/ml-base/lr-bdd-tests/common/utils/token"
	"github.wdf.sap.corp/ml-base/lr-bdd-tests/constants"
	"github.wdf.sap.corp/ml-base/lr-bdd-tests/models/api/request"
	"github.wdf.sap.corp/ml-base/lr-bdd-tests/models/api/response"
)

var _ = Describe("Learning Recommender API Tests: ", func() {

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

		Context("for a single student", func() {
			studID := "00000034"

			reqBody := []requestmodels.RecommendationRequest{
				requestmodels.RecommendationRequest{
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

				resObj := responsemodels.RecommenderResponse{}
				json.Unmarshal(body, &resObj)

				retDataArr := []responsemodels.RestReturnData{}
				json.Unmarshal(resObj.RestOperationStatusVOX.Data.RestReturnData, &retDataArr)

				//Expectations
				Expect(resObj.RestOperationStatusVOX.Status).To(Equal("SUCCESS"))
				Expect(retDataArr[0].Recommendations).NotTo(HaveLen(0))
				Expect(retDataArr[0].StudID).To(Equal(studID))
			})

			Context("with Tenant Header missing", func() {
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

					resObj := responsemodels.RecommenderResponse{}
					json.Unmarshal(body, &resObj)

					var errorString string
					json.Unmarshal(resObj.RestOperationStatusVOX.Data.RestReturnData, &errorString)

					//Expectations
					Expect(res.StatusCode).To(Equal(400))
					Expect(resObj.RestOperationStatusVOX.Status).To(Equal("FAILURE"))
					Expect(errorString).To(Equal("Missing Required details to process this request"))
				})
			})

			Context("with invalid Tenant ID", func() {
				It("should return an error", func() {
					reqJson, _ := json.Marshal(reqBody)
					headers := []rest.Header{
						rest.Header{Key: "authorization", Value: "Bearer " + accessToken},
						rest.Header{Key: "content-type", Value: "application/json"},
						rest.Header{Key: "tenant_id", Value: "invalid_tenant_id"},
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
					Expect(res.StatusCode).To(Equal(400))
					Expect(resObj.RestOperationStatusVOX.Status).To(Equal("FAILURE"))
				})
			})

			Context("with an invalid student id", func() {
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

					resObj := responsemodels.RecommenderResponse{}
					json.Unmarshal(body, &resObj)

					retDataArr := []responsemodels.RestReturnData{}
					json.Unmarshal(resObj.RestOperationStatusVOX.Data.RestReturnData, &retDataArr)

					//Expectations
					Expect(res.StatusCode).To(Equal(200))
					Expect(resObj.RestOperationStatusVOX.Status).To(Equal("SUCCESS"))
					Expect(retDataArr[0].Recommendations).To(HaveLen(0))
					Expect(retDataArr[0].StudID).To(Equal(studID))
				})
			})

		})

		Context("for a batch of students", func() {

			studentIds := []string{"00000034", "00000044"}
			reqBody := []requestmodels.RecommendationRequest{}

			for _, studID := range studentIds {
				reqBody = append(reqBody, requestmodels.RecommendationRequest{
					StudentID: studID,
				})
			}

			Context("having all valid student IDs", func() {
				It("should return recommended courses for all students", func() {
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

					resObj := responsemodels.RecommenderResponse{}
					json.Unmarshal(body, &resObj)

					retDataArr := []responsemodels.RestReturnData{}
					json.Unmarshal(resObj.RestOperationStatusVOX.Data.RestReturnData, &retDataArr)

					//Expectations
					Expect(res.StatusCode).To(Equal(200))
					Expect(resObj.RestOperationStatusVOX.Status).To(Equal("SUCCESS"))
					Expect(retDataArr).To(HaveLen(len(reqBody)))
					Expect(retDataArr[0].Recommendations).NotTo(HaveLen(0))
					Expect(retDataArr[0].StudID).To(Equal(studentIds[0]))
					Expect(retDataArr[1].Recommendations).NotTo(HaveLen(0))
					Expect(retDataArr[1].StudID).To(Equal(studentIds[1]))
				})
			})

			Context("having some invalid student IDs", func() {
				It("should return recommended courses for all valid student IDs", func() {
					//add some invalid ID to existing valid IDs
					reqBody = append(reqBody, requestmodels.RecommendationRequest{
						StudentID: "00000000",
					})
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

					resObj := responsemodels.RecommenderResponse{}
					json.Unmarshal(body, &resObj)

					retDataArr := []responsemodels.RestReturnData{}
					json.Unmarshal(resObj.RestOperationStatusVOX.Data.RestReturnData, &retDataArr)

					//Expectations
					Expect(res.StatusCode).To(Equal(200))
					Expect(resObj.RestOperationStatusVOX.Status).To(Equal("SUCCESS"))
					Expect(retDataArr).To(HaveLen(len(reqBody) - 1))
				})
			})
		})

		Context("with an invalid oauth token", func() {
			It("should return an error", func() {
				reqBody := []requestmodels.RecommendationRequest{
					requestmodels.RecommendationRequest{
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
				resObj := responsemodels.APIHubErrorResponse{}
				json.Unmarshal(body, &resObj)

				//Expectations
				Expect(res.StatusCode).To(Equal(401))
				Expect(resObj.Fault.FaultString).To(Equal("Invalid Access Token"))

			})
		})

		Context("with a specific correlation ID", func() {
			It("should return same correlation ID as part of response", func() {
				reqBody := []requestmodels.RecommendationRequest{
					requestmodels.RecommendationRequest{
						StudentID: "00000034",
					},
				}

				reqJson, _ := json.Marshal(reqBody)

				uuid, _ := gorand.UUID()
				headers := []rest.Header{
					rest.Header{Key: "authorization", Value: "Bearer " + accessToken},
					rest.Header{Key: "content-type", Value: "application/json"},
					rest.Header{Key: "tenant_id", Value: env.Tenant},
					rest.Header{Key: "X_CORRELATION_ID", Value: uuid},
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
				Expect(resObj.RestOperationStatusVOX.CorrelationID).To(Equal(uuid))
			})
		})
	})
})
