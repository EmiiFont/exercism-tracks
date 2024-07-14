package flatten

func Flatten(nested interface{}) []interface{} {
	flattened := []interface{}{}
	for _, v := range nested.([]interface{}) {
		switch v.(type) {
		case []interface{}:
			flattened = append(flattened, Flatten(v)...)
		case int:
			flattened = append(flattened, v)
		}
	}
	return flattened
}
