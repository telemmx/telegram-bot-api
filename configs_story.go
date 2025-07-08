package tgbotapi

/*
postStory
Posts a story on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns Story on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection
content	InputStoryContent	Yes	Content of the story
active_period	Integer	Yes	Period after which the story is moved to the archive, in seconds; must be one of 6 * 3600, 12 * 3600, 86400, or 2 * 86400
caption	String	Optional	Caption of the story, 0-2048 characters after entities parsing
parse_mode	String	Optional	Mode for parsing entities in the story caption. See formatting options for more details.
caption_entities	Array of MessageEntity	Optional	A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
areas	Array of StoryArea	Optional	A JSON-serialized list of clickable areas to be shown on the story
post_to_chat_page	Boolean	Optional	Pass True to keep the story accessible after it expires
protect_content	Boolean	Optional	Pass True if the content of the story must be protected from forwarding and screenshotting
editStory
Edits a story previously posted by the bot on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns Story on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection
story_id	Integer	Yes	Unique identifier of the story to edit
content	InputStoryContent	Yes	Content of the story
caption	String	Optional	Caption of the story, 0-2048 characters after entities parsing
parse_mode	String	Optional	Mode for parsing entities in the story caption. See formatting options for more details.
caption_entities	Array of MessageEntity	Optional	A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
areas	Array of StoryArea	Optional	A JSON-serialized list of clickable areas to be shown on the story
deleteStory
Deletes a story previously posted by the bot on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns True on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection
story_id	Integer	Yes	Unique identifier of the story to delete
*/

type InputStoryContent struct {
	Type  string // photo, video
	Photo string

	Video               string
	Duration            int
	IsAnimated          bool
	CoverFrameTimestamp float32
}

/*
LocationAddress
Describes the physical address of a location.

Field	Type	Description
country_code	String	The two-letter ISO 3166-1 alpha-2 country code of the country where the location is located
state	String	Optional. State of the location
city	String	Optional. City of the location
street	String	Optional. Street address of the location
*/

type LocationAddress struct {
	CountryCode string
	State       string
	City        string
	Street      string
}

/*
StoryAreaPosition
Describes the position of a clickable area within a story.

Field	Type	Description
x_percentage	Float	The abscissa of the area's center, as a percentage of the media width
y_percentage	Float	The ordinate of the area's center, as a percentage of the media height
width_percentage	Float	The width of the area's rectangle, as a percentage of the media width
height_percentage	Float	The height of the area's rectangle, as a percentage of the media height
rotation_angle	Float	The clockwise rotation angle of the rectangle, in degrees; 0-360
corner_radius_percentage	Float	The radius of the rectangle corner rounding, as a percentage of the media width
*/

type StoryAreaPosition struct {
	XPercentage            float64
	YPercentage            float64
	WidthPercentage        float64
	HeightPercentage       float64
	RotationAngle          float64
	CornerRadiusPercentage float64
}

type StoryAreaType struct {
	Type string //location ,suggested_reaction

	Latitude  float64
	Longitude float64
	Address   LocationAddress

	/**
	reaction_type	ReactionType	Type of the reaction
	is_dark	Boolean	Optional. Pass True if the reaction area has a dark background
	is_flipped	Boolean	Optional. Pass True if reaction area corner is flipped
	*/
	ReactionType string
	IsDark       bool
	IsFlipped    bool

	//“link”
	URL string

	// weather
	//temperature	Float	Temperature, in degree Celsius
	// emoji	String	Emoji representing the weather
	// background_color	Integer	A color of the area background in the ARGB format
	Temperature     float64
	Emoji           string
	BackgroundColor int

	//name	String	Unique name of the gift
	Name string
}

type StoryArea struct {
	Type     StoryAreaType
	Position StoryAreaPosition
}

type PostStoryConfig struct {
	BusinessConnectionID string
	Content              InputStoryContent
	ActivePeriod         int
	Caption              string
	ParseMode            string
	CaptionEntities      []MessageEntity
	Areas                []StoryArea
	PostToChatPage       bool
	ProtectContent       bool
}

func (config PostStoryConfig) method() string {
	return "postStory"
}

func (config PostStoryConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddInterface("content", config.Content)
	params.AddNonZero("active_period", config.ActivePeriod)
	params.AddNonEmpty("caption", config.Caption)
	params.AddNonEmpty("parse_mode", config.ParseMode)
	params.AddInterface("caption_entities", config.CaptionEntities)
	params.AddInterface("areas", config.Areas)
	params.AddBool("post_to_chat_page", config.PostToChatPage)
	params.AddBool("protect_content", config.ProtectContent)

	err := params.CheckArgs("business_connection_id", "content", "active_period")
	if err != nil {
		return nil, err
	}
	return params, nil
}

type EditStoryConfig struct {
	BusinessConnectionID string
	StoryID              int
	Content              InputStoryContent
	Caption              string
	ParseMode            string
	CaptionEntities      []MessageEntity
	Areas                []StoryArea
	PostToChatPage       bool
	ProtectContent       bool
}

func (config EditStoryConfig) method() string {
	return "editStory"
}

func (config EditStoryConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddNonZero("story_id", config.StoryID)
	params.AddInterface("content", config.Content)
	params.AddNonEmpty("caption", config.Caption)
	params.AddNonEmpty("parse_mode", config.ParseMode)
	params.AddInterface("caption_entities", config.CaptionEntities)
	params.AddInterface("areas", config.Areas)
	params.AddBool("post_to_chat_page", config.PostToChatPage)
	params.AddBool("protect_content", config.ProtectContent)

	err := params.CheckArgs("business_connection_id", "story_id", "content")
	if err != nil {
		return nil, err
	}
	return params, nil
}

type DeleteStoryConfig struct {
	BusinessConnectionID string
	StoryID              int
}

func (config DeleteStoryConfig) method() string {
	return "deleteStory"
}

func (config DeleteStoryConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddNonZero("story_id", config.StoryID)
	err := params.CheckArgs("business_connection_id", "story_id")
	if err != nil {
		return nil, err
	}
	return params, nil
}
