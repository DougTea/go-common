package web

type SimpleSavedResult struct {
	Saved bool `json:"saved"`
}

type SimpleCreatedResult struct {
	Created bool `json:"created"`
}

type SimpleDeletedResult struct {
	Deleted bool `json:"deleted"`
}

type SimpleUpdatedResult struct {
	Updated bool `json:"updated"`
}
