package internal

type FieldMeta struct {
	Participate bool // indicates if this field participates in all stuff that this lib does

	ValueSources map[string]string // map of sources of value. E.g {"env": "HOST", "json": "host"}
	DefaultValue string            // default value is stored as string because we parse it from string
}
