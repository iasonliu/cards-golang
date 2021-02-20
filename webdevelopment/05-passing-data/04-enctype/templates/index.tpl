{{template "header"}}

<form method="POST" enctype="multipart/form-data">
    <label for="firstName">First Name</label>
    <input type="text" id="firstName" name="first">
    <br>
    <label for="lastName">Last Name</label>
    <input type="text" id="lastName" name="last">
    <br>
    <label for="sub">Subscribe</label>
    <input type="checkbox" id="sub" name="subscribe">
    <br>
    <input type="file" name="filename">
    <br>
    <input type="submit">
</form>

<br>

{{if .}}
<h1>BODY: {{.}}</h1>
{{end}}

{{template "footer"}}