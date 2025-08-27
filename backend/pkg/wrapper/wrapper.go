package wrapper

import "encoding/json"

func Response(status bool, message, data any) string {
	result, err := json.Marshal(map[string]any{
		"status":  status,
		"message": message,
		"data":    data,
	})
	if err != nil {
		panic(err)
	}
	return string(result)
}
