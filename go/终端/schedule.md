# schedule



```
import "atomicgo.dev/schedule"
```



Package schedule provides a simple scheduler for Go.

It can run a function at a given time, in a given duration, or repeatedly at a given interval.

## Index



- type Task
	- [func After(duration time.Duration, task func()) *Task](https://github.com/atomicgo/schedule#After)
	- [func At(t time.Time, task func()) *Task](https://github.com/atomicgo/schedule#At)
	- [func Every(interval time.Duration, task func() bool) *Task](https://github.com/atomicgo/schedule#Every)
	- [func (s *Task) ExecutesIn() time.Duration](https://github.com/atomicgo/schedule#Task.ExecutesIn)
	- [func (s *Task) IsActive() bool](https://github.com/atomicgo/schedule#Task.IsActive)
	- [func (s *Task) NextExecutionTime() time.Time](https://github.com/atomicgo/schedule#Task.NextExecutionTime)
	- [func (s *Task) StartedAt() time.Time](https://github.com/atomicgo/schedule#Task.StartedAt)
	- [func (s *Task) Stop()](https://github.com/atomicgo/schedule#Task.Stop)
	- [func (s *Task) Wait()](https://github.com/atomicgo/schedule#Task.Wait)



## type [Task](https://github.com/atomicgo/schedule/blob/main/schedule.go#L6-L10)



Task holds information about the running task and can be used to stop running tasks.

```
type Task struct {
    // contains filtered or unexported fields
}
```





### func [After](https://github.com/atomicgo/schedule/blob/main/schedule.go#L58)



```
func After(duration time.Duration, task func()) *Task
```



After executes the task after the given duration. The function is non-blocking. If you want to wait for the task to be executed, use the Task.Wait method.

<details style="box-sizing: border-box; display: block; margin-top: 0px; margin-bottom: var(--base-size-16); color: rgb(31, 35, 40); font-family: -apple-system, BlinkMacSystemFont, &quot;Segoe UI&quot;, &quot;Noto Sans&quot;, Helvetica, Arial, sans-serif, &quot;Apple Color Emoji&quot;, &quot;Segoe UI Emoji&quot;; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;"><summary style="box-sizing: border-box; display: list-item; cursor: pointer;">Example</summary><p dir="auto" style="box-sizing: border-box; margin-top: 0px; margin-bottom: var(--base-size-16);"></p><div class="highlight highlight-source-go notranslate position-relative overflow-auto" dir="auto" style="box-sizing: border-box; position: relative !important; overflow: auto !important; margin-bottom: var(--base-size-16); display: flex; justify-content: space-between; background-color: var(--bgColor-muted, var(--color-canvas-subtle));"><pre style="box-sizing: border-box; font-family: var(--fontStack-monospace, ui-monospace, SFMono-Regular, SF Mono, Menlo, Consolas, Liberation Mono, monospace); font-size: 13.6px; margin-top: 0px; margin-bottom: 0px; overflow-wrap: normal; padding: var(--base-size-16); overflow: auto; line-height: 1.45; color: var(--fgColor-default, var(--color-fg-default)); background-color: var(--bgColor-muted, var(--color-canvas-subtle)); border-radius: 6px; word-break: normal; min-height: 52px;"><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span></pre><div class="zeroclipboard-container" style="box-sizing: border-box; display: block; animation: auto ease 0s 1 normal none running none;"><clipboard-copy aria-label="Copy" class="ClipboardButton btn btn-invisible js-clipboard-copy m-2 p-0 d-flex flex-justify-center flex-items-center" data-copy-feedback="Copied!" data-tooltip-direction="w" value="package main

import (
	&quot;fmt&quot;
	&quot;time&quot;

	&quot;atomicgo.dev/schedule&quot;
)

func main() {
	task := schedule.After(5*time.Second, func() {
		fmt.Println(&quot;5 seconds are over!&quot;)
	})

	fmt.Println(&quot;Some stuff happening...&quot;)

	task.Wait()
}" tabindex="0" role="button" style="box-sizing: border-box; position: relative; display: flex !important; padding: 0px !important; font-size: 14px; font-weight: var(--base-text-weight-medium, 500); line-height: 20px; white-space: nowrap; vertical-align: middle; cursor: pointer; user-select: none; border: 0px; border-radius: 6px; appearance: none; color: var(--fgColor-accent, var(--color-accent-fg)); background-color: transparent; box-shadow: none; transition: color 80ms cubic-bezier(0.33, 1, 0.68, 1) 0s, background-color, box-shadow, border-color; justify-content: center !important; align-items: center !important; margin: var(--base-size-8, 8px) !important; width: var(--control-small-size, 28px); height: var(--control-small-size, 28px);"><svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-copy js-clipboard-copy-icon"><path d="M0 6.75C0 5.784.784 5 1.75 5h1.5a.75.75 0 0 1 0 1.5h-1.5a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-1.5a.75.75 0 0 1 1.5 0v1.5A1.75 1.75 0 0 1 9.25 16h-7.5A1.75 1.75 0 0 1 0 14.25Z"></path><path d="M5 1.75C5 .784 5.784 0 6.75 0h7.5C15.216 0 16 .784 16 1.75v7.5A1.75 1.75 0 0 1 14.25 11h-7.5A1.75 1.75 0 0 1 5 9.25Zm1.75-.25a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-7.5a.25.25 0 0 0-.25-.25Z"></path></svg></clipboard-copy></div></div><p dir="auto" style="box-sizing: border-box; margin-top: 0px; margin-bottom: var(--base-size-16);"></p></details>



### func [At](https://github.com/atomicgo/schedule/blob/main/schedule.go#L77)



```
func At(t time.Time, task func()) *Task
```



At executes the task at the given time. The function is non-blocking. If you want to wait for the task to be executed, use the Task.Wait method.

<details style="box-sizing: border-box; display: block; margin-top: 0px; margin-bottom: var(--base-size-16); color: rgb(31, 35, 40); font-family: -apple-system, BlinkMacSystemFont, &quot;Segoe UI&quot;, &quot;Noto Sans&quot;, Helvetica, Arial, sans-serif, &quot;Apple Color Emoji&quot;, &quot;Segoe UI Emoji&quot;; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;"><summary style="box-sizing: border-box; display: list-item; cursor: pointer;">Example</summary><p dir="auto" style="box-sizing: border-box; margin-top: 0px; margin-bottom: var(--base-size-16);"></p><div class="highlight highlight-source-go notranslate position-relative overflow-auto" dir="auto" style="box-sizing: border-box; position: relative !important; overflow: auto !important; margin-bottom: var(--base-size-16); display: flex; justify-content: space-between; background-color: var(--bgColor-muted, var(--color-canvas-subtle));"><pre style="box-sizing: border-box; font-family: var(--fontStack-monospace, ui-monospace, SFMono-Regular, SF Mono, Menlo, Consolas, Liberation Mono, monospace); font-size: 13.6px; margin-top: 0px; margin-bottom: 0px; overflow-wrap: normal; padding: var(--base-size-16); overflow: auto; line-height: 1.45; color: var(--fgColor-default, var(--color-fg-default)); background-color: var(--bgColor-muted, var(--color-canvas-subtle)); border-radius: 6px; word-break: normal; min-height: 52px;"><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span></pre><div class="zeroclipboard-container" style="box-sizing: border-box; display: block; animation: auto ease 0s 1 normal none running none;"><clipboard-copy aria-label="Copy" class="ClipboardButton btn btn-invisible js-clipboard-copy m-2 p-0 d-flex flex-justify-center flex-items-center" data-copy-feedback="Copied!" data-tooltip-direction="w" value="package main

import (
	&quot;fmt&quot;
	&quot;time&quot;

	&quot;atomicgo.dev/schedule&quot;
)

func main() {
	task := schedule.At(time.Now().Add(5*time.Second), func() {
		fmt.Println(&quot;5 seconds are over!&quot;)
	})

	fmt.Println(&quot;Some stuff happening...&quot;)

	task.Wait()
}" tabindex="0" role="button" style="box-sizing: border-box; position: relative; display: flex !important; padding: 0px !important; font-size: 14px; font-weight: var(--base-text-weight-medium, 500); line-height: 20px; white-space: nowrap; vertical-align: middle; cursor: pointer; user-select: none; border: 0px; border-radius: 6px; appearance: none; color: var(--fgColor-accent, var(--color-accent-fg)); background-color: transparent; box-shadow: none; transition: color 80ms cubic-bezier(0.33, 1, 0.68, 1) 0s, background-color, box-shadow, border-color; justify-content: center !important; align-items: center !important; margin: var(--base-size-8, 8px) !important; width: var(--control-small-size, 28px); height: var(--control-small-size, 28px);"><svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-copy js-clipboard-copy-icon"><path d="M0 6.75C0 5.784.784 5 1.75 5h1.5a.75.75 0 0 1 0 1.5h-1.5a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-1.5a.75.75 0 0 1 1.5 0v1.5A1.75 1.75 0 0 1 9.25 16h-7.5A1.75 1.75 0 0 1 0 14.25Z"></path><path d="M5 1.75C5 .784 5.784 0 6.75 0h7.5C15.216 0 16 .784 16 1.75v7.5A1.75 1.75 0 0 1 14.25 11h-7.5A1.75 1.75 0 0 1 5 9.25Zm1.75-.25a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-7.5a.25.25 0 0 0-.25-.25Z"></path></svg></clipboard-copy></div></div><p dir="auto" style="box-sizing: border-box; margin-top: 0px; margin-bottom: var(--base-size-16);"></p></details>



### func [Every](https://github.com/atomicgo/schedule/blob/main/schedule.go#L96)



```
func Every(interval time.Duration, task func() bool) *Task
```



Every executes the task in the given interval, as long as the task function returns true. The function is non-blocking. If you want to wait for the task to be executed, use the Task.Wait method.

<details style="box-sizing: border-box; display: block; margin-top: 0px; margin-bottom: var(--base-size-16); color: rgb(31, 35, 40); font-family: -apple-system, BlinkMacSystemFont, &quot;Segoe UI&quot;, &quot;Noto Sans&quot;, Helvetica, Arial, sans-serif, &quot;Apple Color Emoji&quot;, &quot;Segoe UI Emoji&quot;; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;"><summary style="box-sizing: border-box; display: list-item; cursor: pointer;">Example</summary><p dir="auto" style="box-sizing: border-box; margin-top: 0px; margin-bottom: var(--base-size-16);"></p><div class="highlight highlight-source-go notranslate position-relative overflow-auto" dir="auto" style="box-sizing: border-box; position: relative !important; overflow: auto !important; margin-bottom: var(--base-size-16); display: flex; justify-content: space-between; background-color: var(--bgColor-muted, var(--color-canvas-subtle));"><pre style="box-sizing: border-box; font-family: var(--fontStack-monospace, ui-monospace, SFMono-Regular, SF Mono, Menlo, Consolas, Liberation Mono, monospace); font-size: 13.6px; margin-top: 0px; margin-bottom: 0px; overflow-wrap: normal; padding: var(--base-size-16); overflow: auto; line-height: 1.45; color: var(--fgColor-default, var(--color-fg-default)); background-color: var(--bgColor-muted, var(--color-canvas-subtle)); border-radius: 6px; word-break: normal; min-height: 52px;"><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-smi" style="box-sizing: border-box; color: var(--color-prettylights-syntax-storage-modifier-import);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-k" style="box-sizing: border-box; color: var(--color-prettylights-syntax-keyword);"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-c" style="box-sizing: border-box; color: var(--color-prettylights-syntax-comment);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-s" style="box-sizing: border-box; color: var(--color-prettylights-syntax-string);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-c1" style="box-sizing: border-box; color: var(--color-prettylights-syntax-constant);"></span><span class="pl-s1" style="box-sizing: border-box;"></span><span class="pl-en" style="box-sizing: border-box; color: var(--color-prettylights-syntax-entity);"></span></pre><div class="zeroclipboard-container" style="box-sizing: border-box; display: block; animation: auto ease 0s 1 normal none running none;"><clipboard-copy aria-label="Copy" class="ClipboardButton btn btn-invisible js-clipboard-copy m-2 p-0 d-flex flex-justify-center flex-items-center" data-copy-feedback="Copied!" data-tooltip-direction="w" value="package main

import (
	&quot;fmt&quot;
	&quot;time&quot;

	&quot;atomicgo.dev/schedule&quot;
)

func main() {
	task := schedule.Every(time.Second, func() bool {
		fmt.Println(&quot;1 second is over!&quot;)

		return true // return false to stop the task
	})

	fmt.Println(&quot;Some stuff happening...&quot;)

	time.Sleep(10 * time.Second)

	task.Stop()
}" tabindex="0" role="button" style="box-sizing: border-box; position: relative; display: flex !important; padding: 0px !important; font-size: 14px; font-weight: var(--base-text-weight-medium, 500); line-height: 20px; white-space: nowrap; vertical-align: middle; cursor: pointer; user-select: none; border: 0px; border-radius: 6px; appearance: none; color: var(--fgColor-accent, var(--color-accent-fg)); background-color: transparent; box-shadow: none; transition: color 80ms cubic-bezier(0.33, 1, 0.68, 1) 0s, background-color, box-shadow, border-color; justify-content: center !important; align-items: center !important; margin: var(--base-size-8, 8px) !important; width: var(--control-small-size, 28px); height: var(--control-small-size, 28px);"><svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-copy js-clipboard-copy-icon"><path d="M0 6.75C0 5.784.784 5 1.75 5h1.5a.75.75 0 0 1 0 1.5h-1.5a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-1.5a.75.75 0 0 1 1.5 0v1.5A1.75 1.75 0 0 1 9.25 16h-7.5A1.75 1.75 0 0 1 0 14.25Z"></path><path d="M5 1.75C5 .784 5.784 0 6.75 0h7.5C15.216 0 16 .784 16 1.75v7.5A1.75 1.75 0 0 1 14.25 11h-7.5A1.75 1.75 0 0 1 5 9.25Zm1.75-.25a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-7.5a.25.25 0 0 0-.25-.25Z"></path></svg></clipboard-copy></div></div><p dir="auto" style="box-sizing: border-box; margin-top: 0px; margin-bottom: var(--base-size-16);"></p></details>



### func (*Task) [ExecutesIn](https://github.com/atomicgo/schedule/blob/main/schedule.go#L31)



```
func (s *Task) ExecutesIn() time.Duration
```



ExecutesIn returns the duration until the next execution.



### func (*Task) [IsActive](https://github.com/atomicgo/schedule/blob/main/schedule.go#L36)



```
func (s *Task) IsActive() bool
```



IsActive returns true if the scheduler is active.



### func (*Task) [NextExecutionTime](https://github.com/atomicgo/schedule/blob/main/schedule.go#L26)



```
func (s *Task) NextExecutionTime() time.Time
```



NextExecutionTime returns the time when the next execution will happen.



### func (*Task) [StartedAt](https://github.com/atomicgo/schedule/blob/main/schedule.go#L21)



```
func (s *Task) StartedAt() time.Time
```



StartedAt returns the time when the scheduler was started.



### func (*Task) [Stop](https://github.com/atomicgo/schedule/blob/main/schedule.go#L52)



```
func (s *Task) Stop()
```



Stop stops the scheduler.



### func (*Task) [Wait](https://github.com/atomicgo/schedule/blob/main/schedule.go#L47)



```
func (s *Task) Wait()
```



Wait blocks until the scheduler is stopped. After and At will stop automatically after the task is executed.

Generated by [gomarkdoc](https://github.com/princjef/gomarkdoc)