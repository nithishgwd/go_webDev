package web

const indexTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>{{.Title}}</title>
</head>
<body>
	<form action="/" method="post">
		<label for="first_name">First Name:</label>
		<input type="text" name="first_name" required>
		<br>
		<label for="last_name">Last Name:</label>
		<input type="text" name="last_name" required>
		<br>
		<input type="submit" value="Submit">
	</form>
</body>
</html>
`

const welcomeTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>{{.Title}}</title>
</head>
<body>
	<h1>Hello, {{.Name}}! Welcome to our website.</h1>
</body>
</html>
`
