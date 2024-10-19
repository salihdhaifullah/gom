# Gom package
The`Gom` package provides a set of utilities for generating Markdown formatted text in Go.
It simplifies the process of creating structured documents with various formatting options

---

## Features
- #### Component-Based: *Enables easy composition of Markdown documents using reusable components.*

- #### Headers: *Easily create headers of different levels (H1 to H6).*

- #### Text Formatting: *Apply styles like bold, italic, strikethrough, and inline code.*

- #### Lists: *Generate unordered (UL) and ordered (OL) lists.*

- #### Task Lists: *Create task lists with checkboxes.*

- #### Links and Images: *Include hyperlinks and images in the Markdown output.*

- #### Code Blocks: *Add code snippets with language specification for syntax highlighting.*

- #### Quotes: *Create block quotes for important messages.*

- #### Conditional Rendering: *Render elements based on conditions.*


---

## Functions
#### Escape
Escapes punctuation in the input string so it allows for using special characters like \*\, \~\, \#\, etc\.\.\.
**Example:** 

```go
Escape("*****")
```
#### Doc
Joins multiple nodes into a single node, it works like a fragment
**Example:** 

```go
Doc(H1("Title"), UL("Item 1", "Item 2"))
```
#### Headers
These functions create headers of different levels. For example:
```go
H1("Header 1")
```

```go
H2("Header 2")
```

```go
H3("Header 3")
```

```go
H4("Header 4")
```

```go
H5("Header 5")
```

```go
H6("Header 6")
```
#### Text Formatting
You can format text in italics or bold using these functions:
**Italic Example:** 

```go
Italic("This is italic text")
```
**Bold Example:** 

```go
Bold("This is bold text")
```
**Using Both:** 

```go
Bold("This is bold text and", Italic("this is the italic part"))
```
#### Lists
Generate lists using the following functions:
**Unordered List Example:** 

```go
UL("Item 1", "Item 2")
```
**Ordered List Example:** 

```go
OL("First", "Second")
```
#### Quotes and Task Lists
**Block Quote Example:** 

```go
Quote("This is a block quote.")
```
**Task List Example:** 

```go
Task(false, "This is a task item that is not done.")
```

```go
Task(true, "This is a task item that is done.")
```
#### Links and Images
**Link Example:** 

```go
Link("https://example.com", "This is a link")
```
**Image Example:** 

```go
Img("https://example.com/image.png", "This is an image")
```
#### Code
**Code Block Example:** 

```go
CodeBlock("go", "fmt.Println('Hello, World!')")
```
#### Inline Code

```go
Code("fmt.Println")
```
#### Conditionals
You can conditionally render nodes based on a boolean value:
**If Example:** 

```go
If(condition, "This will be shown")
```
**If else Example:** 

```go
IfElse(condition, "Shown", "Not Shown")
```
#### Iteration
You can iterate over a list and apply a component to each item:
**For Example:** 

```go

data := []string{"item one", "item two", "item three"}
Component := func (item string) string {
    return Doc(
        H1("this is the component"),
        Quote(Escape(item)),
    )
}

For(data, Component)
```


---

## Full Code Example

```go

package main

import (
	"fmt"
	"math/rand"
	"time"
        . "github.com/salihdhaifullah/gom"
)

// PageTitle component that generates a unique page title with dynamic time
func PageTitle() string {
	return Doc(
		H1("Welcome to Your Personalized Page"),
		"The content below is dynamically generated based on current state.", L,
		"Page generated at: ", Bold(time.Now().Format("Mon, 02 Jan 2006 15:04:05")), L,
	)
}

// RandomFact component generates a random fact each time it's rendered
func RandomFact() string {
	facts := []string{
		"Bananas are berries, but strawberries are not.",
		"Honey never spoils.",
		"A group of flamingos is called a 'flamboyance'.",
		"Octopuses have three hearts.",
		"Sharks existed before trees.",
	}

	randomIndex := rand.Intn(len(facts))

	return Doc(
		H2("Did You Know?"),
		Quote(facts[randomIndex]), L,
	)
}

// UserMessage component that takes a name and a dynamic message as input
func UserMessage(name string) string {
	return Doc(
		H2("Personalized Message"),
		"Hello, ", Bold(name), "!", L,
		"We hope you're enjoying your day. Here's something special just for you:", L,
		Bold("Keep being awesome and stay curious!"), L,
	)
}

// TodoList component dynamically generates a task list
type Todo struct {
	done bool
	text string
}

func TodoList(tasks []Todo) string {
	return Doc(
		H2("Your To-Do List"),
		For(tasks, func(todo Todo) string { return Task(todo.done, todo.text) }),
	)
}

// Footer component that shows a dynamically updated year
func Footer() string {
	return Doc(
		H3("Footer"),
		"Thank you for visiting this page. ", L,
		"© ", fmt.Sprintf("%d", time.Now().Year()), " All rights reserved.",
	)
}

// MainPage function that renders all components together
func MainPage(name string, todos []Todo) string {
	return Doc(
		PageTitle(),
		RandomFact(),
		UserMessage(name),
		TodoList(todos),
		Footer(),
	)
}

func GetTodos() []Todo {
	return []Todo{
		{
			text: "Learn Go",
			done: true,
		},

		{
			text: "Contribute to open source",
			done: false,
		},

		{
			text: "Read about dynamic programming",
			done: false,
		},
	}

}

func main() {
	tasks := GetTodos()
	fmt.Println(MainPage("Alice", tasks))
}

```
