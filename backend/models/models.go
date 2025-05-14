package models

import "time"

type Victima struct {
    ID            int       `json:"id"`
    FullName      string    `json:"full_name"`
    Cause         string    `json:"cause,omitempty"`
    Details       string    `json:"details,omitempty"`
    CreatedAt     time.Time `json:"created_at"`
    DeathTime     time.Time `json:"death_time"`
    ImageURL      string    `json:"image_url"`
    CauseAdded    bool      `json:"cause_added"`
    DetailsAdded  bool      `json:"details_added"`
    IsDead        bool      `json:"is_dead"`
}