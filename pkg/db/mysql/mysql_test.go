package mysql

import (
	"github.com/polaris-team/test-report-server/pkg/config"
	"github.com/polaris-team/test-report-server/pkg/utils/json"
	"testing"

	"strconv"
)

func TestAvailability(t *testing.T) {
	t.Logf("start load config")
	config.LoadConfig("F:\\workspace-golang-polaris\\polaris-backend\\polaris-testing\\configs", "application")

	t.Logf(json.ToJson(config.GetConfig()))
	conn, err := GetConnect()
	if err != nil {
		t.Log(err)
	}
	t.Log(settings)
	count, _ := conn.Collection("ppm_tst_plan").Find().Count()
	t.Logf(strconv.FormatUint(count, 10))
}
