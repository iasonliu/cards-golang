<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>INDEX</title>
</head>
<body>

<h1> COOKIE VALUES</h1>
{{range .}}
<h2> {{.}} </h2>
{{end}}

<form method="post" enctype="multipart/form-data">
<input type="file" name="filename" >
<input type="submit">
</form>
</body>
</html>