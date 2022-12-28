package translations

import "taiga-mm/config"

type TranslationString int

const (
	TaskCreated TranslationString = iota
	TaskDeleted
	SubtaskCreated
	UnknownAction
	FullMessage
	CommentDeleted
	CommentChanged
	CommentAdded
	TaskRenamed
	TaskDescriptionChanged
	TaskAsigneeChanged
	TaskAsigneeAdded
	TaskAsigneeDeleted
	SubtaskAsigneeChanged
	SubtaskAsigneeAdded
	SubtaskAsigneeDeleted
	StatusChanged
	TagDeleted
	TagAdded
	DueDateChanged
	DueDateAdded
	DueDateDeleted
	AttachmentAdded
	AttachmentDeleted
	AttachmentChanged
	TeamRequirementOn
	TeamRequirementOff
	ClientRequirementOn
	ClientRequirementOff
	TaskBlocked
	TaskUnblocked
	TaskChanged
)

var Translations = map[string]map[TranslationString]string{
	"ru": RussianTranslation,
	"en": EnglishTranslation,
	// "template": TemplateTranslation,
}

func Get(s TranslationString) string {
	return Translations[config.Config.Language][s]
}
