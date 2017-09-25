package lr_bdd_test

import (
	"flag"
	"net/http"

	"github.com/coding-yogi/go_bdd/handlers"
	"github.com/coding-yogi/go_bdd/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var (
	environment string
	client      *http.Client
	env         models.Environment
)

func init() {
	flag.StringVar(&environment, "environment", "qa", "environment is used to set execution env")
}
func TestGoBdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoBdd Suite")
}

var _ = BeforeSuite(func() {
	client = &http.Client{}
	env, _ = config.GetEnvDetails("qa")
})
