package requests

type CreateDiaryRequest struct {
    Date        string   `json:"date" `
    Mood        string   `json:"mood"`
    Emotions    []string `json:"emotions"`
    Description string   `json:"description"`
    UserID      string   `json:"user_id"`
}

type UpdateDiaryRequest struct{
    Mood        string   `json:"mood"`
    Emotions    []string `json:"emotions"`
    Description string   `json:"description"`
}