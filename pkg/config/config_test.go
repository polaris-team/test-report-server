package config

import (
	"github.com/polaris-team/test-report-server/pkg/utils/json"
	"testing"
)

func TestCombination(t *testing.T) {
	LoadConfig("F:\\workspace-golang-polaris\\polaris-backend\\polaris-testing\\configs", "application")

	bs, _ := json.ToJson(conf)
	t.Log(string(bs))
}
