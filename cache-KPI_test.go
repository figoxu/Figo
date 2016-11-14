package Figo

import "testing"

func TestCachKpi(t *testing.T) {
	kpi := &CacheKPI{}
	kpi.Save(true)
	kpi.Save(false)
	kpi.Save(true)
	kpi.Save(true)
	kpi.Print()
}
