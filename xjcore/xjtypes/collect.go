package xjtypes

import (
	"reflect"
)

func CollectString(list interface{}, field string) []string {
	var vlas []string
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		len := reflect.ValueOf(list).Len()
		for i := 0; i < len; i++ {
			v := reflect.ValueOf(list).Index(i)
			ve := v.Interface()
			if reflect.TypeOf(ve).Kind() == reflect.Struct {
				for j := 0; j < reflect.TypeOf(ve).NumField(); j++ {
					name := reflect.TypeOf(ve).Field(j).Name
					if name == field {
						val := reflect.ValueOf(ve).Field(j).String()
						vlas = append(vlas, val)
					}
				}
			}
		}
	}
	return vlas
}

func CollectInt(list interface{}, field string) []int {
	var vlas []int
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		len := reflect.ValueOf(list).Len()
		for i := 0; i < len; i++ {
			v := reflect.ValueOf(list).Index(i)
			ve := v.Interface()
			if reflect.TypeOf(ve).Kind() == reflect.Struct {
				for j := 0; j < reflect.TypeOf(ve).NumField(); j++ {
					name := reflect.TypeOf(ve).Field(j).Name
					if name == field {
						val := reflect.ValueOf(ve).Field(j).Int()
						vlas = append(vlas, int(val))
					}
				}
			}
		}
	}
	return vlas
}

func CollectInt64(list interface{}, field string) []int64 {
	var vlas []int64
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		len := reflect.ValueOf(list).Len()
		for i := 0; i < len; i++ {
			v := reflect.ValueOf(list).Index(i)
			ve := v.Interface()
			if reflect.TypeOf(ve).Kind() == reflect.Struct {
				for j := 0; j < reflect.TypeOf(ve).NumField(); j++ {
					name := reflect.TypeOf(ve).Field(j).Name
					if name == field {
						val := reflect.ValueOf(ve).Field(j).Int()
						vlas = append(vlas, val)
					}
				}
			}
		}
	}
	return vlas
}

func CollectFloat64(list interface{}, field string) []float64 {
	var vlas []float64
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		len := reflect.ValueOf(list).Len()
		for i := 0; i < len; i++ {
			v := reflect.ValueOf(list).Index(i)
			ve := v.Interface()
			if reflect.TypeOf(ve).Kind() == reflect.Struct {
				for j := 0; j < reflect.TypeOf(ve).NumField(); j++ {
					name := reflect.TypeOf(ve).Field(j).Name
					if name == field {
						val := reflect.ValueOf(ve).Field(j).Float()
						vlas = append(vlas, val)
					}
				}
			}
		}
	}
	return vlas
}

func CollectSetString(list interface{}, field string) []string {
	var vals []string
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		len := reflect.ValueOf(list).Len()
		for i := 0; i < len; i++ {
			v := reflect.ValueOf(list).Index(i)
			ve := v.Interface()
			if reflect.TypeOf(ve).Kind() == reflect.Struct {
				for j := 0; j < reflect.TypeOf(ve).NumField(); j++ {
					name := reflect.TypeOf(ve).Field(j).Name
					if name == field {
						val := reflect.ValueOf(ve).Field(j).String()
						find := false
						for _, item := range vals {
							if val == item {
								find = true
								break
							}
						}
						if !find {
							vals = append(vals, val)
						}
					}
				}
			}
		}
	}
	return vals
}

func CollectSetInt(list interface{}, field string) []int {
	var vals []int
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		len := reflect.ValueOf(list).Len()
		for i := 0; i < len; i++ {
			v := reflect.ValueOf(list).Index(i)
			ve := v.Interface()
			if reflect.TypeOf(ve).Kind() == reflect.Struct {
				for j := 0; j < reflect.TypeOf(ve).NumField(); j++ {
					name := reflect.TypeOf(ve).Field(j).Name
					if name == field {
						val := reflect.ValueOf(ve).Field(j).Int()
						find := false
						for _, item := range vals {
							if int(val) == item {
								find = true
								break
							}
						}
						if !find {
							vals = append(vals, int(val))
						}
					}
				}
			}
		}
	}
	return vals
}

func CollectSetInt64(list interface{}, field string) []int64 {
	var vals []int64
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		len := reflect.ValueOf(list).Len()
		for i := 0; i < len; i++ {
			v := reflect.ValueOf(list).Index(i)
			ve := v.Interface()
			if reflect.TypeOf(ve).Kind() == reflect.Struct {
				for j := 0; j < reflect.TypeOf(ve).NumField(); j++ {
					name := reflect.TypeOf(ve).Field(j).Name
					if name == field {
						val := reflect.ValueOf(ve).Field(j).Int()
						find := false
						for _, item := range vals {
							if val == item {
								find = true
								break
							}
						}
						if !find {
							vals = append(vals, val)
						}
					}
				}
			}
		}
	}
	return vals
}

func CollectSetFloat64(list interface{}, field string) []float64 {
	var vals []float64
	if reflect.TypeOf(list).Kind() == reflect.Slice {
		len := reflect.ValueOf(list).Len()
		for i := 0; i < len; i++ {
			v := reflect.ValueOf(list).Index(i)
			ve := v.Interface()
			if reflect.TypeOf(ve).Kind() == reflect.Struct {
				for j := 0; j < reflect.TypeOf(ve).NumField(); j++ {
					name := reflect.TypeOf(ve).Field(j).Name
					if name == field {
						val := reflect.ValueOf(ve).Field(j).Float()
						find := false
						for _, item := range vals {
							if val == item {
								find = true
								break
							}
						}
						if !find {
							vals = append(vals, val)
						}
					}
				}
			}
		}
	}
	return vals
}
