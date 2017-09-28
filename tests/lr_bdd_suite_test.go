package lr_bdd_test

import (
	"flag"
	"net/http"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"github.wdf.sap.corp/ml-base/lr-bdd/handlers"
	"github.wdf.sap.corp/ml-base/lr-bdd/models/appconfig"

	"testing"
)

var (
	environment string
	client      *http.Client
	env         appconfig.Environment
)

func init() {
	flag.StringVar(&environment, "environment", "qa", "environment is used to set execution env")
}
func TestGoBdd(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Learning Recommender API Test Suite", []Reporter{junitReporter})
}

var _ = BeforeSuite(func() {
	client = &http.Client{}
	env, _ = config.GetEnvDetails("qa")
})
