package static

import "embed"

//go:embed styles/*
var StyleFiles embed.FS

//go:embed media/*
var ImageFiles embed.FS
