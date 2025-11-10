package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Format(data []byte, formatType string) (string, error) {
	switch formatType {
	case "string":
		return fmt.Sprintf("%s", data), nil
	case "json":
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, data, "", "	")
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s", prettyJSON.Bytes()), nil
	case "raw":
		return fmt.Sprint(data), nil
	}

	return "", fmt.Errorf("unknown format type: %s", formatType)
}
