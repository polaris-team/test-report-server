package copyer

import "github.com/polaris-team/test-report-server/pkg/utils/json"

func Copy(src interface{}, source interface{})error{
	jsonStr, err := json.ToJson(src)
	if err != nil{
		return err
	}
	json.FromJson(jsonStr, source)
	return nil
}
