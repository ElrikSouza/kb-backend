package session

type SessionStore interface {
	// saves the payload on the session store, and returns an opaque token identifying that session
	SaveSession(payload SessionPayload) (string, error)

	// uses the opaque token to get the session payload
	RetrieveSession(sessionToken string) (SessionPayload, error)
}
