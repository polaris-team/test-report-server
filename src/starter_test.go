package main

import "testing"

func TestGenerateStruct(t *testing.T) {
	err := GenerateStruct("../graphql/models", "ppm_pro_project")
	t.Log(err)
}