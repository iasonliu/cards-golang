<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Hello World!</title>
</head>
<body>
<ul>
    {{range $index, $element := .}}
    <li>{{$index}} - {{$element}}</li>
    {{end}}
</ul>
</body>
</html>