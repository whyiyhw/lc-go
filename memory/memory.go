package memory

type InputValues map[string]any

func getInputValue(values InputValues, key string) (any, error) {
	if key == "" {
		return nil, nil
	}

	value, ok := values[key]
	if !ok {
		return nil, nil
	}

	return value, nil
}
