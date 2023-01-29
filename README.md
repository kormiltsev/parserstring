# Find on page

Package returns string between 2 preset uniq strings

## Usage
- Example 1

```
page := "package main
func myFunctionName(parameter string) {
...
}
func myOtherFunc() {
...
}"

	request := parserstring.NewReq()
	request.AddRequestToken("function name", "func ", "(")
	request.ParseString(page)
```
Returns `[map[function name:myFunctionName] map[function name:myOtherFunc]]`

- Example 2
```
page := "<thead><tr><th>Pollutant</th></tr></thead><tbody><tr><td>Unhealthy for Sensitive Groups</td><td> 114 <span class=\"sett\">AQI level</span><app-trend..."

	request := parserstring.NewReq()
	request.AddRequestToken("Status", "<tbody><tr><td>", "</td>")
	request.AddRequestToken("Index", "</td><td> ", " <span")
	request.ParseString(page)
```
Returns `[map[Status:Unhealthy for Sensitive Groups Index:114]]`

## ✨Features

- skips token if not found
- tokens order is important
- search for 1<sup>st</sup> token starts from begining, but next tokens searching starts from place where previous token was found
- searching repeats till end of string given

