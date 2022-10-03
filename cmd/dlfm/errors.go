package main

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrNilDGoSession = Error("nil discordgo session")
)