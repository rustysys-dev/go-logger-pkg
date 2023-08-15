package scheme

type Logger interface {
	// // WithContext embeds the context into the logger
	// WithContext(context.Context) Logger
	// // ToContext inserts the logger into a given context
	// ToContext(context.Context) context.Context
	// // FromContext retrieves the logger from a given context
	// FromContext(context.Context) Logger
	// Adds Item to the internal logger context
	AddItem(string, any) Logger
	// Debug writes a message with any context
	Debug(string, ...any)
	// // Debug writes a message with any context
	// Info(string, ...any)
	// // Debug writes a message with any context
	// Warning(string, ...any)
	// // Debug writes a message with any context
	// Error(string, ...any)
	// // Debug writes a message with any context
	// Fatal(string, ...any)
}
