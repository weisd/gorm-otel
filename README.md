# gorm-otel

opentelemetry support for gorm2.

### Features

- [x] Record `SQL` in `span` logs.
- [x] Record `Result` in `span` logs.
- [x] Record `Table` in `span` tags.
- [x] Record `Error` in `span` tags and logs.
- [x] Register `Create` `Query` `Delete` `Update` `Row` `Raw` tracing callbacks. 

### Get Started

I assume that you already have an opentelemetry Tracer client started in your project.

```go
func main() {
	var db *gorm.DB
	
	db.Use(gormotel.New())
	
	// if you want to use customized tracer instead of opentelemetry.GlobalTracer() which is default,
	// you can use the option WithTracer(yourTracer)
}
```
