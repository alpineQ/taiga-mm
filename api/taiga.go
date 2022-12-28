package api

// Data represents the data sent in the request from Taiga
type TaigaRequest struct {
	Type   string `json:"type"`
	Date   string `json:"date"`
	Action string `json:"action"`
	Data   Data   `json:"data"`
	Change Change `json:"change"`
	By     User   `json:"by"`
}

type Data struct {
	CustomAttributesValues map[string]string `json:"custom_attributes_values,omitempty"`
	ID                     int64             `json:"id"`
	Ref                    int64             `json:"ref,omitempty"`
	Watchers               []int             `json:"watchers,omitempty"`
	Permalink              string            `json:"permalink"`
	Tags                   []string          `json:"tags,omitempty"`
	Project                Project           `json:"project"`
	Owner                  User              `json:"owner"`
	AssignedTo             User              `json:"assigned_to,omitempty"`
	AssignedUsers          []int64           `json:"assigned_users,omitempty"`
	Status                 Status            `json:"status"`
	UserStory              UserStory         `json:"user_story,omitempty"`
	Subject                string            `json:"subject"`
	Milestone              Milestone         `json:"milestone,omitempty"`
	By                     User              `json:"by,omitempty"`
	ExternalReference      string            `json:"external_reference"`
	Points                 []Point           `json:"points"`
	IsBlocked              bool              `json:"is_blocked"`
	BlockedNote            string            `json:"blocked_note"`
	IsClosed               bool              `json:"is_closed"`
	CreatedDate            string            `json:"created_date"`
	ModifiedDate           string            `json:"modified_date"`
	FinishDate             string            `json:"finish_date"`
	DueDate                string            `json:"due_date"`
	DueDateReason          string            `json:"due_date_reason"`
	Description            string            `json:"description"`
	ClientRequirement      bool              `json:"client_requirement"`
	TeamRequirement        bool              `json:"team_requirement"`
	GeneratedFromIssue     string            `json:"generated_from_issue"`
	GeneratedFromTask      string            `json:"generated_from_task"`
	TribeGig               string            `json:"tribe_gig"`
}

type Project struct {
	ID        int64  `json:"id"`
	Permalink string `json:"permalink"`
	Name      string `json:"name"`
	Logo      string `json:"logo_big_url,omitempty"`
}

type UserStory struct {
	CustomAttributesValues map[string]string `json:"custom_attributes_values"`
	Watchers               []int             `json:"watchers"`
	Permalink              string            `json:"permalink"`
	Tags                   []string          `json:"tags"`
	ExternalReference      string            `json:"external_reference"`
	Project                Project           `json:"project"`
	Owner                  User              `json:"owner"`
	AssignedTo             User              `json:"assigned_to"`
	Points                 []Point           `json:"points"`
	Status                 Status            `json:"status"`
	Milestone              Milestone         `json:"milestone"`
	ID                     int64             `json:"id"`
	IsBlocked              bool              `json:"is_blocked"`
	BlockedNote            string            `json:"blocked_note"`
	Ref                    int64             `json:"ref"`
	IsClosed               bool              `json:"is_closed"`
	CreatedDate            string            `json:"created_date"`
	ModifiedDate           string            `json:"modified_date"`
	FinishDate             string            `json:"finish_date"`
	Subject                string            `json:"subject"`
	Description            string            `json:"description"`
	ClientRequirement      bool              `json:"client_requirement"`
	TeamRequirement        bool              `json:"team_requirement"`
	GeneratedFromIssue     string            `json:"generated_from_issue"`
	TribeGig               string            `json:"tribe_gig"`
}

func (story *UserStory) Equal(other *UserStory) bool {
	return story.ID == other.ID && story.Permalink == other.Permalink
}

type Change struct {
	Diff              Diff             `json:"diff"`
	Comment           string           `json:"comment,omitempty"`
	CommentHTML       string           `json:"comment_html,omitempty"`
	CommentVersions   []CommentVersion `json:"comment_versions,omitempty"`
	DeleteCommentDate string           `json:"delete_comment_date,omitempty"`
	EditCommentDate   string           `json:"edit_comment_date,omitempty"`
}

type User struct {
	ID         int64  `json:"id"`
	Permalink  string `json:"permalink"`
	Username   string `json:"username"`
	FullName   string `json:"full_name"`
	Photo      string `json:"photo,omitempty"`
	GravatarID string `json:"gravatar_id"`
}

type Status struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Color      string `json:"color"`
	IsClosed   bool   `json:"is_closed"`
	IsArchived bool   `json:"is_archived"`
}

type Point struct {
	Role  string  `json:"role"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type Milestone struct {
	Permalink       string  `json:"permalink"`
	Project         Project `json:"project"`
	Owner           User    `json:"owner"`
	ID              int64   `json:"id"`
	Name            string  `json:"name"`
	Slug            string  `json:"slug"`
	EstimatedStart  string  `json:"estimated_start"`
	EstimatedFinish string  `json:"estimated_finish"`
	CreatedDate     string  `json:"created_date"`
	ModifiedDate    string  `json:"modified_date"`
	Closed          bool    `json:"closed"`
	Disponibility   float64 `json:"disponibility"`
}

type Diff struct {
	AssignedTo        StringDiff  `json:"assigned_to,omitempty"`
	AssignedUsers     StringDiff  `json:"assigned_users,omitempty"`
	DescriptionDiff   string      `json:"description_diff,omitempty"`
	Status            StringDiff  `json:"status,omitempty"`
	Tags              ListDiff    `json:"tags,omitempty"`
	DueDate           StringDiff  `json:"due_date,omitempty"`
	Subject           StringDiff  `json:"subject,omitempty"`
	ContentHTML       StringDiff  `json:"content_html,omitempty"`
	ContentDiff       StringDiff  `json:"content_diff,omitempty"`
	Attachments       Attachments `json:"attachments,omitempty"`
	TeamRequirement   BoolDiff    `json:"team_requirement,omitempty"`
	ClientRequirement BoolDiff    `json:"client_requirement,omitempty"`
	IsBlocked         BoolDiff    `json:"is_blocked,omitempty"`
	BlockedNoteDiff   StringDiff  `json:"blocked_note_diff,omitempty"`
	BlockedNoteHTML   StringDiff  `json:"blocked_note_html,omitempty"`
}

type BoolDiff struct {
	From bool `json:"from"`
	To   bool `json:"to"`
}

type StringDiff struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type ListDiff struct {
	From []string `json:"from"`
	To   []string `json:"to"`
}

func (diff *ListDiff) GetDiff() string {
	var biggerList []string
	var smallerList []string
	if len(diff.From) > len(diff.To) {
		biggerList = diff.From
		smallerList = diff.To
	} else if len(diff.From) < len(diff.To) {
		biggerList = diff.To
		smallerList = diff.From
	} else {
		return ""
	}
	for _, tagA := range biggerList {
		found := false
		for _, tagB := range smallerList {
			if tagA == tagB {
				found = true
				break
			}
		}
		if !found {
			return tagA
		}
	}
	return ""
}

type Attachments struct {
	New     []Attachment `json:"new"`
	Changed []Attachment `json:"changed"`
	Deleted []Attachment `json:"deleted"`
}

type Attachment struct {
	Filename string            `json:"filename"`
	URL      string            `json:"url"`
	ThumbURL string            `json:"thumb_url"`
	Changes  AttachmentChanges `json:"changes"`
}

type AttachmentChanges struct {
	Description []string `json:"description"`
}

type CommentVersion struct {
	Date        string `json:"date"`
	User        UserID `json:"user"`
	Comment     string `json:"comment"`
	CommentHTML string `json:"comment_html"`
}

type UserID struct {
	ID int64 `json:"id"`
}
