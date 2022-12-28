package translations

var TemplateTranslation = map[TranslationString]string{
	TaskCreated:            "created a task",
	TaskDeleted:            "deleted a task",
	SubtaskCreated:         "created for a task [%s](%s) new subtask",             // task name, task URL
	UnknownAction:          "done something (%s)",                                 // unknown action
	FullMessage:            "%s has %s %s in %s project",                          // author, action, task, project
	CommentDeleted:         "deleted a comment `\"%s\"` for a task",               // comment
	CommentChanged:         "changed a comment to `\"%s\"` for a task",            // comment
	CommentAdded:           "added a comment `\"%s\"` for a task",                 // comment
	TaskRenamed:            "renamed a task from `%s` to",                         // task name
	TaskDescriptionChanged: "changed the description to `\"%s\"` for a task",      // task description
	TaskAsigneeChanged:     "changed the asignee from %s to %s for a task",        // previous asignee, current asignee
	TaskAsigneeAdded:       "set %s as asignee for a task",                        // asignee
	TaskAsigneeDeleted:     "deleted %s as asignee for a task",                    // previous asignee
	SubtaskAsigneeChanged:  "changed asignee from %s to %s for a subtask",         // previous asignee, current asignee
	SubtaskAsigneeAdded:    "set %s as asignee for a subtask",                     // asignee
	SubtaskAsigneeDeleted:  "deleted %s as asignee for a subtask",                 // previous asignee
	StatusChanged:          "changed status from `\"%s\"` to `\"%s\"` for a task", // previous status, current status
	TagDeleted:             "deleted tag `\"%s\"` for a task",                     // deleted tag
	TagAdded:               "added tag `\"%s\"` for a task",                       // tag
	DueDateChanged:         "changed due date from `%s` to `%s`%s for a task",     // previous due date, current due date, reason to change due date (if present)
	DueDateAdded:           "set due date to `%s`%s for a task",                   // due date, due date reason (if present)
	DueDateDeleted:         "deleted дедлайн на `%s` for a task",                  // previous due date
	AttachmentAdded:        "added [attachment](%s) for a task",                   // attachment URL
	AttachmentChanged:      "changed [attachment](%s) for a task",                 // attachment URL
	AttachmentDeleted:      "deleted attachment for a task",
	TeamRequirementOff:     "disabled a team requirement for a task",
	TeamRequirementOn:      "enabled a team requirement for a task",
	ClientRequirementOff:   "disabled a team requirement for a task",
	ClientRequirementOn:    "enabled a team requirement for a task",
	TaskUnblocked:          "unblocked a task",
	TaskBlocked:            "blocked a task",
}
