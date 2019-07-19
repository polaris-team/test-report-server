package mysql

import (
	"testing"

	"github.com/polaris-team/polaris-backend/polaris-common/core/utils/json"

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
