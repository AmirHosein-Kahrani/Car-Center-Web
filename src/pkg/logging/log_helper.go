package logging

func logParamsToZapParams(extra map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0)
	for k, v := range extra {

		params = append(params, string(k))
		params = append(params, v)

	}
	return params
}

func logParamsToZeroParams(extra map[ExtraKey]interface{}) map[string]interface{} {
	params := map[string]interface{}{}

	for k, v := range extra {
		params[string(k)] = v
	}
	return params
}
