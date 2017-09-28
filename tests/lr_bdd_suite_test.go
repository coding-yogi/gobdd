package lr_bdd_test

import (
	"flag"
	"net/http"

	"github.com/coding-yogi/go_bdd/handlers"
	"github.com/coding-yogi/go_bdd/models/appconfig"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

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
