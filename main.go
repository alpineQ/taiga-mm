package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"taiga-mm/api"
	"taiga-mm/config"
	"taiga-mm/translations"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		fmt.Printf("Unable to parse configuration: %s\n", err)
		return
	}
	if cfg.MattermostToken == "" {
		fmt.Println("mattermost_token variable in config.json is required")
		return
	}
	if _, ok := translations.Translations[cfg.Language]; !ok {
		fmt.Printf(`Your language("%s") is not supported.\nUse "en" or add your translation to the project via pull request.\n`, cfg.Language)
		return
	}
	hostURL := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	http.HandleFunc("/taiga-integration/channel/", integrationHandler)
	fmt.Printf("Listening at http://%s\n", hostURL)
	err = http.ListenAndServe(hostURL, nil)
	if err != nil {
		fmt.Printf("Error while running server: %s\n", err)
		return
	}
}

func integrationHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	data := &api.TaigaRequest{}
	if err := json.Unmarshal(body, data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Printf("unable to unmarshal JSON body: %s\n", err.Error())
		return
	}

	var action string
	switch data.Action {
	case "create":
		emptyStory := api.UserStory{}
		if data.Data.UserStory.Equal(&emptyStory) {
			action = translations.Get(translations.TaskCreated)
		} else {
			subject := data.Data.UserStory.Subject
			permalink := data.Data.UserStory.Permalink
			action = fmt.Sprintf(translations.Get(translations.SubtaskCreated), subject, permalink)
		}
	case "delete":
		action = translations.Get(translations.TaskDeleted)
	case "change":
		action = changeAction(data)
	case "test":
		fmt.Fprint(w, "Test successfull")
		return
	default:
		action = fmt.Sprintf(translations.Get(translations.UnknownAction), data.Action)
	}

	author := fmt.Sprintf("[%s](%s)", data.By.FullName, data.By.Permalink)
	if mmUsername, ok := config.Config.Usernames[data.By.Username]; ok {
		author += fmt.Sprintf("(@%s)", mmUsername)
	}
	project := fmt.Sprintf("[%s](%s)", data.Data.Project.Name, data.Data.Project.Permalink)
	subject := fmt.Sprintf("[%s](%s)", data.Data.Subject, data.Data.Permalink)
	message := fmt.Sprintf(translations.Get(translations.FullMessage), author, action, subject, project)
	channel := r.URL.Path[len("/taiga-integration/channel/"):]
	if err := api.SendMessage(channel, message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Success")
}

func addLink(user string) string {
	if mmUsername, ok := config.Config.Usernames[user]; ok {
		return fmt.Sprintf("%s(@%s)", user, mmUsername)
	}
	return user
}

func changeAction(request *api.TaigaRequest) string {
	data := &request.Change
	if data.Comment != "" {
		if data.DeleteCommentDate != "" {
			return fmt.Sprintf(translations.Get(translations.CommentDeleted), data.Comment)
		}
		if len(data.CommentVersions) > 0 {
			return fmt.Sprintf(translations.Get(translations.CommentChanged), data.Comment)
		}
		return fmt.Sprintf(translations.Get(translations.CommentAdded), data.Comment)
	}
	if data.Diff.Subject.From != "" && data.Diff.Subject.To != "" {
		return fmt.Sprintf(translations.Get(translations.TaskRenamed), data.Diff.Subject.From)
	}
	if data.Diff.DescriptionDiff != "" {
		return fmt.Sprintf(translations.Get(translations.TaskRenamed), request.Data.Description)
	}
	if data.Diff.AssignedUsers.From != "" && data.Diff.AssignedUsers.To != "" {
		return fmt.Sprintf(translations.Get(translations.TaskAsigneeChanged), addLink(data.Diff.AssignedUsers.From), addLink(data.Diff.AssignedUsers.To))
	}
	if data.Diff.AssignedUsers.To != "" {
		return fmt.Sprintf(translations.Get(translations.TaskAsigneeAdded), addLink(data.Diff.AssignedUsers.To))
	}
	if data.Diff.AssignedUsers.From != "" {
		return fmt.Sprintf(translations.Get(translations.TaskAsigneeDeleted), addLink(data.Diff.AssignedUsers.From))
	}
	if data.Diff.AssignedTo.From != "" && data.Diff.AssignedTo.To != "" {
		return fmt.Sprintf(translations.Get(translations.SubtaskAsigneeChanged), addLink(data.Diff.AssignedTo.From), addLink(data.Diff.AssignedTo.To))
	}
	if data.Diff.AssignedTo.To != "" {
		return fmt.Sprintf(translations.Get(translations.SubtaskAsigneeAdded), addLink(data.Diff.AssignedTo.To))
	}
	if data.Diff.AssignedTo.From != "" {
		return fmt.Sprintf(translations.Get(translations.SubtaskAsigneeDeleted), addLink(data.Diff.AssignedTo.From))
	}
	if data.Diff.Status.From != "" && data.Diff.Status.To != "" {
		return fmt.Sprintf(translations.Get(translations.StatusChanged), data.Diff.Status.From, data.Diff.Status.To)
	}
	if len(data.Diff.Tags.From) > len(data.Diff.Tags.To) {
		return fmt.Sprintf(translations.Get(translations.TagDeleted), data.Diff.Tags.GetDiff())
	}
	if len(data.Diff.Tags.From) < len(data.Diff.Tags.To) {
		return fmt.Sprintf(translations.Get(translations.TagAdded), data.Diff.Tags.GetDiff())
	}
	if data.Diff.DueDate.From != "" && data.Diff.DueDate.To != "" {
		dueDateReason := ""
		if request.Data.DueDateReason != "" {
			dueDateReason = fmt.Sprintf(" (%s)", request.Data.DueDateReason)
		}
		return fmt.Sprintf(translations.Get(translations.DueDateChanged), data.Diff.DueDate.From, data.Diff.DueDate.To, dueDateReason)
	}
	if data.Diff.DueDate.To != "" {
		dueDateReason := ""
		if request.Data.DueDateReason != "" {
			dueDateReason = fmt.Sprintf(" (%s)", request.Data.DueDateReason)
		}
		return fmt.Sprintf(translations.Get(translations.DueDateAdded), data.Diff.DueDate.To, dueDateReason)
	}
	if data.Diff.DueDate.From != "" {
		return fmt.Sprintf(translations.Get(translations.DueDateDeleted), data.Diff.DueDate.From)
	}
	if len(data.Diff.Attachments.New) > 0 {
		return fmt.Sprintf(translations.Get(translations.AttachmentAdded), data.Diff.Attachments.New[0].URL)
	}
	if len(data.Diff.Attachments.Deleted) > 0 {
		return translations.Get(translations.AttachmentDeleted)
	}
	if len(data.Diff.Attachments.Changed) > 0 {
		return fmt.Sprintf(translations.Get(translations.AttachmentChanged), data.Diff.Attachments.Changed[0].URL)
	}
	if data.Diff.TeamRequirement.From && !data.Diff.TeamRequirement.To {
		return translations.Get(translations.TeamRequirementOff)
	}
	if data.Diff.TeamRequirement.To && !data.Diff.TeamRequirement.From {
		return translations.Get(translations.TeamRequirementOn)
	}
	if data.Diff.ClientRequirement.From && !data.Diff.ClientRequirement.To {
		return translations.Get(translations.ClientRequirementOff)
	}
	if data.Diff.ClientRequirement.To && !data.Diff.ClientRequirement.From {
		return translations.Get(translations.ClientRequirementOn)
	}
	if data.Diff.IsBlocked.From && !data.Diff.IsBlocked.To {
		return translations.Get(translations.TaskUnblocked)
	}
	if data.Diff.IsBlocked.To && !data.Diff.IsBlocked.From {
		return translations.Get(translations.TaskBlocked)
	}
	return translations.Get(translations.TaskChanged)
}
