<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
</head>
<body>

{{if .First}}
USER NAME {{.UserName}}<br>
PASSWORD {{.Password}}<br>
ROLE {{.Role}}<br>
FIRST {{.First}}<br>
LAST {{.Last}}<br>
{{else}}
<h2><a href="/signup">sign up</a></h2>
<h2><a href="/login">login</a></h2>
{{end}}

<br>
<h2>Go to <a href="/bar">the bar</a></h2>
</body>
</html>