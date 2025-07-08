package tgbotapi

/*
Checklist
Describes a checklist.

Field	Type	Description
title	String	Title of the checklist
title_entities	Array of MessageEntity	Optional. Special entities that appear in the checklist title
tasks	Array of ChecklistTask	List of tasks in the checklist
others_can_add_tasks	True	Optional. True, if users other than the creator of the list can add tasks to the list
others_can_mark_tasks_as_done	True	Optional. True, if users other than the creator of the list can mark tasks as done or not done
*/
type Checklist struct {
	Title                    string          `json:"title"`
	TitleEntities            []MessageEntity `json:"title_entities"`
	Tasks                    []ChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool            `json:"others_can_add_tasks"`
	OthersCanMarkTasksAsDone bool            `json:"others_can_mark_tasks_as_done"`
}

/*
InputChecklist
Describes a checklist to create.

Field	Type	Description
title	String	Title of the checklist; 1-255 characters after entities parsing
parse_mode	String	Optional. Mode for parsing entities in the title. See formatting options for more details.
title_entities	Array of MessageEntity	Optional. List of special entities that appear in the title, which can be specified instead of parse_mode. Currently, only bold, italic, underline, strikethrough, spoiler, and custom_emoji entities are allowed.
tasks	Array of InputChecklistTask	List of 1-30 tasks in the checklist
others_can_add_tasks	Boolean	Optional. Pass True if other users can add tasks to the checklist
others_can_mark_tasks_as_done	Boolean	Optional. Pass True if other users can mark tasks as done or not done in the checklist
*/

type InputChecklist struct {
	Title                    string               `json:"title"`
	ParseMode                string               `json:"parse_mode"`
	TitleEntities            []MessageEntity      `json:"title_entities"`
	Tasks                    []InputChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool                 `json:"others_can_add_tasks"`
	OthersCanMarkTasksAsDone bool                 `json:"others_can_mark_tasks_as_done"`
}

/*
ChecklistTask
Describes a task in a checklist.

Field	Type	Description
id	Integer	Unique identifier of the task
text	String	Text of the task
text_entities	Array of MessageEntity	Optional. Special entities that appear in the task text
completed_by_user	User	Optional. User that completed the task; omitted if the task wasn't completed
completion_date	Integer	Optional. Point in time (Unix timestamp) when the task was completed; 0 if the task wasn't completed
*/
type ChecklistTask struct {
	Id              int             `json:"id"`
	Text            string          `json:"text"`
	TextEntities    []MessageEntity `json:"text_entities"`
	CompletedByUser User            `json:"completed_by_user"`
	CompletionDate  int             `json:"completion_date"`
}

/*
InputChecklistTask
Describes a task to add to a checklist.

Field	Type	Description
id	Integer	Unique identifier of the task; must be positive and unique among all task identifiers currently present in the checklist
text	String	Text of the task; 1-100 characters after entities parsing
parse_mode	String	Optional. Mode for parsing entities in the text. See formatting options for more details.
text_entities	Array of MessageEntity	Optional. List of special entities that appear in the text, which can be specified instead of parse_mode. Currently, only bold, italic, underline, strikethrough, spoiler, and custom_emoji entities are allowed.
*/

type InputChecklistTask struct {
	Id           int             `json:"id"`
	Text         string          `json:"text"`
	ParseMode    string          `json:"parse_mode"`
	TextEntities []MessageEntity `json:"text_entities"`
}

/*
ChecklistTasksDone
Describes a service message about checklist tasks marked as done or not done.

Field	Type	Description
checklist_message	Message	Optional. Message containing the checklist whose tasks were marked as done or not done. Note that the Message object in this field will not contain the reply_to_message field even if it itself is a reply.
marked_as_done_task_ids	Array of Integer	Optional. Identifiers of the tasks that were marked as done
marked_as_not_done_task_ids	Array of Integer	Optional. Identifiers of the tasks that were marked as not done
*/
type ChecklistTasksDone struct {
	ChecklistMessage       Message `json:"checklist_message"`
	MarkedAsDoneTaskIds    []int   `json:"marked_as_done_task_ids"`
	MarkedAsNotDoneTaskIds []int   `json:"marked_as_not_done_task_ids"`
}

/*
ChecklistTasksAdded
Describes a service message about tasks added to a checklist.

Field	Type	Description
checklist_message	Message	Optional. Message containing the checklist to which the tasks were added. Note that the Message object in this field will not contain the reply_to_message field even if it itself is a reply.
tasks	Array of ChecklistTask	List of tasks added to the checklist
*/
type ChecklistTasksAdded struct {
	// Optional. Message containing the checklist to which the tasks were added. Note that the Message object in this field will not contain the reply_to_message field even if it itself is a reply.
	ChecklistMessage Message `json:"checklist_message"`
	// List of tasks added to the checklist
	Tasks ChecklistTask `json:"tasks"`
}
