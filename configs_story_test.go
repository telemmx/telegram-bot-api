package tgbotapi

import (
	"testing"
)

func TestPostStoryConfig(t *testing.T) {
	// Test method name
	config := PostStoryConfig{
		BusinessConnectionID: "test_connection",
		Content: InputStoryContent{
			Type:  "photo",
			Photo: "test_photo_id",
		},
		ActivePeriod: 86400,
		Caption:      "Test story caption",
		ParseMode:    "Markdown",
		CaptionEntities: []MessageEntity{
			{
				Type:   "bold",
				Offset: 0,
				Length: 4,
			},
		},
		Areas: []StoryArea{
			{
				Type: StoryAreaType{
					Type:      "location",
					Latitude:  37.7749,
					Longitude: -122.4194,
					Address: LocationAddress{
						CountryCode: "US",
						State:       "California",
						City:        "San Francisco",
						Street:      "Market St",
					},
				},
				Position: StoryAreaPosition{
					XPercentage:            50.0,
					YPercentage:            50.0,
					WidthPercentage:        20.0,
					HeightPercentage:       20.0,
					RotationAngle:          0.0,
					CornerRadiusPercentage: 5.0,
				},
			},
		},
		PostToChatPage: true,
		ProtectContent: false,
	}

	// Test method name
	if config.method() != "postStory" {
		t.Errorf("Expected method 'postStory', got '%s'", config.method())
	}

	// Test params
	params, err := config.params()
	if err != nil {
		t.Fatalf("Error generating params: %v", err)
	}

	// Check that all fields are properly set in params
	if params["business_connection_id"] != "test_connection" {
		t.Errorf("Expected business_connection_id to be 'test_connection', got '%s'", params["business_connection_id"])
	}

	if params["active_period"] != "86400" {
		t.Errorf("Expected active_period to be '86400', got '%s'", params["active_period"])
	}

	if params["caption"] != "Test story caption" {
		t.Errorf("Expected caption to be 'Test story caption', got '%s'", params["caption"])
	}

	if params["parse_mode"] != "Markdown" {
		t.Errorf("Expected parse_mode to be 'Markdown', got '%s'", params["parse_mode"])
	}

	if params["post_to_chat_page"] != "true" {
		t.Errorf("Expected post_to_chat_page to be 'true', got '%s'", params["post_to_chat_page"])
	}

	if params["protect_content"] != "false" {
		t.Errorf("Expected protect_content to be 'false', got '%s'", params["protect_content"])
	}
}

func TestEditStoryConfig(t *testing.T) {
	// Test method name
	config := EditStoryConfig{
		BusinessConnectionID: "test_connection",
		StoryID:              12345,
		Content: InputStoryContent{
			Type:                "video",
			Video:               "test_video_id",
			Duration:            30,
			IsAnimated:          true,
			CoverFrameTimestamp: 5.5,
		},
		Caption:   "Updated story caption",
		ParseMode: "HTML",
		CaptionEntities: []MessageEntity{
			{
				Type:   "italic",
				Offset: 0,
				Length: 7,
			},
		},
		Areas: []StoryArea{
			{
				Type: StoryAreaType{
					Type:         "suggested_reaction",
					ReactionType: "👍",
					IsDark:       true,
					IsFlipped:    false,
				},
				Position: StoryAreaPosition{
					XPercentage:            30.0,
					YPercentage:            70.0,
					WidthPercentage:        15.0,
					HeightPercentage:       15.0,
					RotationAngle:          45.0,
					CornerRadiusPercentage: 10.0,
				},
			},
		},
		PostToChatPage: false,
		ProtectContent: true,
	}

	// Test method name
	if config.method() != "editStory" {
		t.Errorf("Expected method 'editStory', got '%s'", config.method())
	}

	// Test params
	params, err := config.params()
	if err != nil {
		t.Fatalf("Error generating params: %v", err)
	}

	// Check that all fields are properly set in params
	if params["business_connection_id"] != "test_connection" {
		t.Errorf("Expected business_connection_id to be 'test_connection', got '%s'", params["business_connection_id"])
	}

	if params["story_id"] != "12345" {
		t.Errorf("Expected story_id to be '12345', got '%s'", params["story_id"])
	}

	if params["caption"] != "Updated story caption" {
		t.Errorf("Expected caption to be 'Updated story caption', got '%s'", params["caption"])
	}

	if params["parse_mode"] != "HTML" {
		t.Errorf("Expected parse_mode to be 'HTML', got '%s'", params["parse_mode"])
	}

	if params["post_to_chat_page"] != "false" {
		t.Errorf("Expected post_to_chat_page to be 'false', got '%s'", params["post_to_chat_page"])
	}

	if params["protect_content"] != "true" {
		t.Errorf("Expected protect_content to be 'true', got '%s'", params["protect_content"])
	}
}

func TestDeleteStoryConfig(t *testing.T) {
	// Test method name
	config := DeleteStoryConfig{
		BusinessConnectionID: "test_connection",
		StoryID:              67890,
	}

	// Test method name
	if config.method() != "deleteStory" {
		t.Errorf("Expected method 'deleteStory', got '%s'", config.method())
	}

	// Test params
	params, err := config.params()
	if err != nil {
		t.Fatalf("Error generating params: %v", err)
	}

	// Check that all fields are properly set in params
	if params["business_connection_id"] != "test_connection" {
		t.Errorf("Expected business_connection_id to be 'test_connection', got '%s'", params["business_connection_id"])
	}

	if params["story_id"] != "67890" {
		t.Errorf("Expected story_id to be '67890', got '%s'", params["story_id"])
	}
}
