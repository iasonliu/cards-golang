<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Peeps</title>
</head>
<body>
<ul>
    {{ range .}}
    <li>{{.Name}} -  {{.Age}}</li>
    {{end}}
</ul>
</body>
</html>