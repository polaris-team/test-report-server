package redis

import (
	"testing"

	"github.com/polaris-team/polaris-backend/polaris-common/core/config"
)

func TestAvailability(t *testing.T) {
	t.Logf("start load config")
	config.LoadConfig("F:\\workspace-golang-polaris\\polaris-backend\\polaris-testing\\configs", "application")

	rp := GetProxy()

	rp.SetEx("name", "nico", 60)
	res, _ := rp.Get("name")
	t.Logf(res)
}
