<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Hello World!</title>
</head>
<body>
<ul>
    {{range $key, $val := .}}
    <li>{{$key}} - {{$val}}</li>
    {{end}}
</ul>
</body>
</html>