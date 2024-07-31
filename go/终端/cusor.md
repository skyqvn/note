# cursor



```
import "atomicgo.dev/cursor"
```



Package cursor contains cross-platform methods to move the terminal cursor in different directions. This package can be used to create interactive CLI tools and games, live charts, algorithm visualizations and other updatable output of any kind.

Works niceley with https://github.com/atomicgo/keyboard

Special thanks to github.com/k0kubun/go-ansi which this project is based on.

## Index



- [func Bottom()](https://github.com/atomicgo/cursor#Bottom)
- [func Clear()](https://github.com/atomicgo/cursor#Clear)
- [func ClearLine()](https://github.com/atomicgo/cursor#ClearLine)
- [func ClearLinesDown(n int)](https://github.com/atomicgo/cursor#ClearLinesDown)
- [func ClearLinesUp(n int)](https://github.com/atomicgo/cursor#ClearLinesUp)
- [func Down(n int)](https://github.com/atomicgo/cursor#Down)
- [func DownAndClear(n int)](https://github.com/atomicgo/cursor#DownAndClear)
- [func Hide()](https://github.com/atomicgo/cursor#Hide)
- [func HorizontalAbsolute(n int)](https://github.com/atomicgo/cursor#HorizontalAbsolute)
- [func Left(n int)](https://github.com/atomicgo/cursor#Left)
- [func Move(x, y int)](https://github.com/atomicgo/cursor#Move)
- [func Right(n int)](https://github.com/atomicgo/cursor#Right)
- [func SetTarget(w Writer)](https://github.com/atomicgo/cursor#SetTarget)
- [func Show()](https://github.com/atomicgo/cursor#Show)
- [func StartOfLine()](https://github.com/atomicgo/cursor#StartOfLine)
- [func StartOfLineDown(n int)](https://github.com/atomicgo/cursor#StartOfLineDown)
- [func StartOfLineUp(n int)](https://github.com/atomicgo/cursor#StartOfLineUp)
- [func TestCustomIOWriter(t *testing.T)](https://github.com/atomicgo/cursor#TestCustomIOWriter)
- [func Up(n int)](https://github.com/atomicgo/cursor#Up)
- [func UpAndClear(n int)](https://github.com/atomicgo/cursor#UpAndClear)
- type Area
	- [func NewArea() Area](https://github.com/atomicgo/cursor#NewArea)
	- [func (area *Area) Bottom()](https://github.com/atomicgo/cursor#Area.Bottom)
	- [func (area *Area) Clear()](https://github.com/atomicgo/cursor#Area.Clear)
	- [func (area *Area) ClearLinesDown(n int)](https://github.com/atomicgo/cursor#Area.ClearLinesDown)
	- [func (area *Area) ClearLinesUp(n int)](https://github.com/atomicgo/cursor#Area.ClearLinesUp)
	- [func (area *Area) Down(n int)](https://github.com/atomicgo/cursor#Area.Down)
	- [func (area *Area) DownAndClear(n int)](https://github.com/atomicgo/cursor#Area.DownAndClear)
	- [func (area *Area) Move(x, y int)](https://github.com/atomicgo/cursor#Area.Move)
	- [func (area *Area) StartOfLine()](https://github.com/atomicgo/cursor#Area.StartOfLine)
	- [func (area *Area) StartOfLineDown(n int)](https://github.com/atomicgo/cursor#Area.StartOfLineDown)
	- [func (area *Area) StartOfLineUp(n int)](https://github.com/atomicgo/cursor#Area.StartOfLineUp)
	- [func (area *Area) Top()](https://github.com/atomicgo/cursor#Area.Top)
	- [func (area *Area) Up(n int)](https://github.com/atomicgo/cursor#Area.Up)
	- [func (area *Area) UpAndClear(n int)](https://github.com/atomicgo/cursor#Area.UpAndClear)
	- [func (area *Area) Update(content string)](https://github.com/atomicgo/cursor#Area.Update)
	- [func (area Area) WithWriter(writer Writer) Area](https://github.com/atomicgo/cursor#Area.WithWriter)
- type Cursor
	- [func NewCursor() *Cursor](https://github.com/atomicgo/cursor#NewCursor)
	- [func (c *Cursor) Clear()](https://github.com/atomicgo/cursor#Cursor.Clear)
	- [func (c *Cursor) ClearLine()](https://github.com/atomicgo/cursor#Cursor.ClearLine)
	- [func (c *Cursor) Down(n int)](https://github.com/atomicgo/cursor#Cursor.Down)
	- [func (c *Cursor) Hide()](https://github.com/atomicgo/cursor#Cursor.Hide)
	- [func (c *Cursor) HorizontalAbsolute(n int)](https://github.com/atomicgo/cursor#Cursor.HorizontalAbsolute)
	- [func (c *Cursor) Left(n int)](https://github.com/atomicgo/cursor#Cursor.Left)
	- [func (c *Cursor) Right(n int)](https://github.com/atomicgo/cursor#Cursor.Right)
	- [func (c *Cursor) Show()](https://github.com/atomicgo/cursor#Cursor.Show)
	- [func (c *Cursor) Up(n int)](https://github.com/atomicgo/cursor#Cursor.Up)
	- [func (c *Cursor) WithWriter(w Writer) *Cursor](https://github.com/atomicgo/cursor#Cursor.WithWriter)
- [type Writer](https://github.com/atomicgo/cursor#Writer)



## func [Bottom](https://github.com/atomicgo/cursor/blob/main/utils.go#L86)



```
func Bottom()
```



Bottom moves the cursor to the bottom of the terminal. This is done by calculating how many lines were moved by Up and Down.



## func [Clear](https://github.com/atomicgo/cursor/blob/main/utils.go#L80)



```
func Clear()
```



Clear clears the current position and moves the cursor to the left.



## func [ClearLine](https://github.com/atomicgo/cursor/blob/main/utils.go#L75)



```
func ClearLine()
```



ClearLine clears the current line and moves the cursor to it's start position.



## func [ClearLinesDown](https://github.com/atomicgo/cursor/blob/main/utils.go#L149)



```
func ClearLinesDown(n int)
```



ClearLinesDown clears n lines downwards from the current position and moves the cursor.



## func [ClearLinesUp](https://github.com/atomicgo/cursor/blob/main/utils.go#L142)



```
func ClearLinesUp(n int)
```



ClearLinesUp clears n lines upwards from the current position and moves the cursor.



## func [Down](https://github.com/atomicgo/cursor/blob/main/utils.go#L36)



```
func Down(n int)
```



Down moves the cursor n lines down relative to the current position.



## func [DownAndClear](https://github.com/atomicgo/cursor/blob/main/utils.go#L119)



```
func DownAndClear(n int)
```



DownAndClear moves the cursor down by n lines, then clears the line.



## func [Hide](https://github.com/atomicgo/cursor/blob/main/utils.go#L70)



```
func Hide()
```



Hide the cursor. Don't forget to show the cursor at least at the end of your application with Show. Otherwise the user might have a terminal with a permanently hidden cursor, until they reopen the terminal.



## func [HorizontalAbsolute](https://github.com/atomicgo/cursor/blob/main/utils.go#L56)



```
func HorizontalAbsolute(n int)
```



HorizontalAbsolute moves the cursor to n horizontally. The position n is absolute to the start of the line.



## func [Left](https://github.com/atomicgo/cursor/blob/main/utils.go#L50)



```
func Left(n int)
```



Left moves the cursor n characters to the left relative to the current position.



## func [Move](https://github.com/atomicgo/cursor/blob/main/utils.go#L125)



```
func Move(x, y int)
```



Move moves the cursor relative by x and y.



## func [Right](https://github.com/atomicgo/cursor/blob/main/utils.go#L45)



```
func Right(n int)
```



Right moves the cursor n characters to the right relative to the current position.



## func [SetTarget](https://github.com/atomicgo/cursor/blob/main/utils.go#L25)



```
func SetTarget(w Writer)
```



SetTarget sets to output target of the default curser to the provided cursor.Writer (wrapping io.Writer).



## func [Show](https://github.com/atomicgo/cursor/blob/main/utils.go#L63)



```
func Show()
```



Show the cursor if it was hidden previously. Don't forget to show the cursor at least at the end of your application. Otherwise the user might have a terminal with a permanently hidden cursor, until they reopen the terminal.



## func [StartOfLine](https://github.com/atomicgo/cursor/blob/main/utils.go#L96)



```
func StartOfLine()
```



StartOfLine moves the cursor to the start of the current line.



## func [StartOfLineDown](https://github.com/atomicgo/cursor/blob/main/utils.go#L101)



```
func StartOfLineDown(n int)
```



StartOfLineDown moves the cursor down by n lines, then moves to cursor to the start of the line.



## func [StartOfLineUp](https://github.com/atomicgo/cursor/blob/main/utils.go#L107)



```
func StartOfLineUp(n int)
```



StartOfLineUp moves the cursor up by n lines, then moves to cursor to the start of the line.



## func [TestCustomIOWriter](https://github.com/atomicgo/cursor/blob/main/cursor_test_linux.go#L9)



```
func TestCustomIOWriter(t *testing.T)
```



TestCustomIOWriter tests the cursor functions with a custom Writer.



## func [Up](https://github.com/atomicgo/cursor/blob/main/utils.go#L30)



```
func Up(n int)
```



Up moves the cursor n lines up relative to the current position.



## func [UpAndClear](https://github.com/atomicgo/cursor/blob/main/utils.go#L113)



```
func UpAndClear(n int)
```



UpAndClear moves the cursor up by n lines, then clears the line.



## type [Area](https://github.com/atomicgo/cursor/blob/main/area.go#L10-L15)



Area displays content which can be updated on the fly. You can use this to create live output, charts, dropdowns, etc.

```
type Area struct {
    // contains filtered or unexported fields
}
```





### func [NewArea](https://github.com/atomicgo/cursor/blob/main/area.go#L18)



```
func NewArea() Area
```



NewArea returns a new Area.



### func (*Area) [Bottom](https://github.com/atomicgo/cursor/blob/main/area.go#L86)



```
func (area *Area) Bottom()
```



Bottom moves the cursor to the bottom of the terminal. This is done by calculating how many lines were moved by Up and Down.



### func (*Area) [Clear](https://github.com/atomicgo/cursor/blob/main/area.go#L36)



```
func (area *Area) Clear()
```



Clear clears the content of the Area.



### func (*Area) [ClearLinesDown](https://github.com/atomicgo/cursor/blob/main/area.go#L157)



```
func (area *Area) ClearLinesDown(n int)
```



ClearLinesDown clears n lines downwards from the current position and moves the cursor.



### func (*Area) [ClearLinesUp](https://github.com/atomicgo/cursor/blob/main/area.go#L147)



```
func (area *Area) ClearLinesUp(n int)
```



ClearLinesUp clears n lines upwards from the current position and moves the cursor.



### func (*Area) [Down](https://github.com/atomicgo/cursor/blob/main/area.go#L73)



```
func (area *Area) Down(n int)
```



Down moves the cursor of the area down one line.



### func (*Area) [DownAndClear](https://github.com/atomicgo/cursor/blob/main/area.go#L126)



```
func (area *Area) DownAndClear(n int)
```



DownAndClear moves the cursor down by n lines, then clears the line.



### func (*Area) [Move](https://github.com/atomicgo/cursor/blob/main/area.go#L132)



```
func (area *Area) Move(x, y int)
```



Move moves the cursor relative by x and y.



### func (*Area) [StartOfLine](https://github.com/atomicgo/cursor/blob/main/area.go#L103)



```
func (area *Area) StartOfLine()
```



StartOfLine moves the cursor to the start of the current line.



### func (*Area) [StartOfLineDown](https://github.com/atomicgo/cursor/blob/main/area.go#L108)



```
func (area *Area) StartOfLineDown(n int)
```



StartOfLineDown moves the cursor down by n lines, then moves to cursor to the start of the line.



### func (*Area) [StartOfLineUp](https://github.com/atomicgo/cursor/blob/main/area.go#L114)



```
func (area *Area) StartOfLineUp(n int)
```



StartOfLineUp moves the cursor up by n lines, then moves to cursor to the start of the line.



### func (*Area) [Top](https://github.com/atomicgo/cursor/blob/main/area.go#L95)



```
func (area *Area) Top()
```



Top moves the cursor to the top of the area. This is done by calculating how many lines were moved by Up and Down.



### func (*Area) [Up](https://github.com/atomicgo/cursor/blob/main/area.go#L61)



```
func (area *Area) Up(n int)
```



Up moves the cursor of the area up one line.



### func (*Area) [UpAndClear](https://github.com/atomicgo/cursor/blob/main/area.go#L120)



```
func (area *Area) UpAndClear(n int)
```



UpAndClear moves the cursor up by n lines, then clears the line.



### func (*Area) [Update](https://github.com/atomicgo/cursor/blob/main/area.go#L53)



```
func (area *Area) Update(content string)
```



Update overwrites the content of the Area and adjusts its height based on content.



### func (Area) [WithWriter](https://github.com/atomicgo/cursor/blob/main/area.go#L28)



```
func (area Area) WithWriter(writer Writer) Area
```



WithWriter sets the custom writer.



## type [Cursor](https://github.com/atomicgo/cursor/blob/main/cursor.go#L9-L11)



Cursor displays content which can be updated on the fly. You can use this to create live output, charts, dropdowns, etc.

```
type Cursor struct {
    // contains filtered or unexported fields
}
```





### func [NewCursor](https://github.com/atomicgo/cursor/blob/main/cursor.go#L14)



```
func NewCursor() *Cursor
```



NewCursor creates a new Cursor instance writing to os.Stdout.



### func (*Cursor) [Clear](https://github.com/atomicgo/cursor/blob/main/cursor_other.go#L65)



```
func (c *Cursor) Clear()
```



Clear clears the current position and moves the cursor to the left.



### func (*Cursor) [ClearLine](https://github.com/atomicgo/cursor/blob/main/cursor_other.go#L60)



```
func (c *Cursor) ClearLine()
```



ClearLine clears the current line and moves the cursor to it's start position.



### func (*Cursor) [Down](https://github.com/atomicgo/cursor/blob/main/cursor_other.go#L18)



```
func (c *Cursor) Down(n int)
```



Down moves the cursor n lines down relative to the current position.



### func (*Cursor) [Hide](https://github.com/atomicgo/cursor/blob/main/cursor_other.go#L55)



```
func (c *Cursor) Hide()
```



Hide the cursor. Don't forget to show the cursor at least at the end of your application with Show. Otherwise the user might have a terminal with a permanently hidden cursor, until they reopen the terminal.



### func (*Cursor) [HorizontalAbsolute](https://github.com/atomicgo/cursor/blob/main/cursor_other.go#L40)



```
func (c *Cursor) HorizontalAbsolute(n int)
```



HorizontalAbsolute moves the cursor to n horizontally. The position n is absolute to the start of the line.



### func (*Cursor) [Left](https://github.com/atomicgo/cursor/blob/main/cursor_other.go#L32)



```
func (c *Cursor) Left(n int)
```



Left moves the cursor n characters to the left relative to the current position.



### func (*Cursor) [Right](https://github.com/atomicgo/cursor/blob/main/cursor_other.go#L25)



```
func (c *Cursor) Right(n int)
```



Right moves the cursor n characters to the right relative to the current position.



### func (*Cursor) [Show](https://github.com/atomicgo/cursor/blob/main/cursor_other.go#L48)



```
func (c *Cursor) Show()
```



Show the cursor if it was hidden previously. Don't forget to show the cursor at least at the end of your application. Otherwise the user might have a terminal with a permanently hidden cursor, until they reopen the terminal.



### func (*Cursor) [Up](https://github.com/atomicgo/cursor/blob/main/cursor_other.go#L11)



```
func (c *Cursor) Up(n int)
```



Up moves the cursor n lines up relative to the current position.



### func (*Cursor) [WithWriter](https://github.com/atomicgo/cursor/blob/main/cursor.go#L20)



```
func (c *Cursor) WithWriter(w Writer) *Cursor
```



WithWriter allows for any arbitrary Writer to be used for cursor movement abstracted.



## type [Writer](https://github.com/atomicgo/cursor/blob/main/utils.go#L18-L21)



Writer is an expanded io.Writer interface with a file descriptor.

```
type Writer interface {
    io.Writer
    Fd() uintptr
}
```



Generated by [gomarkdoc](https://github.com/princjef/gomarkdoc)