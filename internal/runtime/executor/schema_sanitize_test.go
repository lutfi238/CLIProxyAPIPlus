package executor

import (
	"encoding/json"
	"testing"
)

func TestSanitizeToolParametersSchema_NestedArrayMissingItems(t *testing.T) {
	// Reproduce the exact error pattern from the user's report:
	// properties.config.items.properties.containerNodes is type: "array" but missing "items"
	input := `{
		"type": "object",
		"properties": {
			"config": {
				"type": "array",
				"items": {
					"type": "object",
					"properties": {
						"containerNodes": {
							"type": "array"
						}
					}
				}
			}
		}
	}`

	result := sanitizeToolParametersSchema(input)

	// Parse the result and check that containerNodes now has "items"
	var schema map[string]interface{}
	if err := json.Unmarshal([]byte(result), &schema); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	// Navigate: properties -> config -> items -> properties -> containerNodes
	props := schema["properties"].(map[string]interface{})
	config := props["config"].(map[string]interface{})
	items := config["items"].(map[string]interface{})
	innerProps := items["properties"].(map[string]interface{})
	containerNodes := innerProps["containerNodes"].(map[string]interface{})

	if _, hasItems := containerNodes["items"]; !hasItems {
		t.Errorf("containerNodes should have 'items' after sanitization, but it doesn't. Got: %+v", containerNodes)
	}

	t.Logf("Sanitized schema: %s", result)
}

func TestSanitizeToolParametersSchema_AlreadyValid(t *testing.T) {
	input := `{"type":"object","properties":{"name":{"type":"string"}}}`
	result := sanitizeToolParametersSchema(input)

	if result != input {
		t.Errorf("expected unchanged schema for valid input\ngot:  %s\nwant: %s", result, input)
	}
}

func TestSanitizeToolParametersSchema_TopLevelArrayMissingItems(t *testing.T) {
	input := `{"type":"array"}`
	result := sanitizeToolParametersSchema(input)

	var schema map[string]interface{}
	if err := json.Unmarshal([]byte(result), &schema); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	if _, hasItems := schema["items"]; !hasItems {
		t.Errorf("expected 'items' to be added")
	}
}

func TestNormalizeGitHubCopilotChatTools_SanitizesNestedSchema(t *testing.T) {
	// Full tool payload in OpenAI chat completions format with a bad nested schema
	body := []byte(`{
		"model": "gpt-5.4",
		"messages": [],
		"tools": [
			{
				"type": "function",
				"function": {
					"name": "mcp__pencil__spawn_agents",
					"description": "Spawn agents",
					"parameters": {
						"type": "object",
						"properties": {
							"config": {
								"type": "array",
								"items": {
									"type": "object",
									"properties": {
										"containerNodes": {
											"type": "array"
										}
									}
								}
							}
						}
					}
				}
			}
		]
	}`)

	result := normalizeGitHubCopilotChatTools(body)

	var payload map[string]interface{}
	if err := json.Unmarshal(result, &payload); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	tools := payload["tools"].([]interface{})
	if len(tools) != 1 {
		t.Fatalf("expected 1 tool, got %d", len(tools))
	}

	tool := tools[0].(map[string]interface{})
	fn := tool["function"].(map[string]interface{})
	params := fn["parameters"].(map[string]interface{})
	props := params["properties"].(map[string]interface{})
	config := props["config"].(map[string]interface{})
	items := config["items"].(map[string]interface{})
	innerProps := items["properties"].(map[string]interface{})
	containerNodes := innerProps["containerNodes"].(map[string]interface{})

	if _, hasItems := containerNodes["items"]; !hasItems {
		t.Errorf("containerNodes should have 'items' after normalization, but it doesn't. Got: %+v", containerNodes)
	}

	t.Logf("Result body: %s", string(result))
}
