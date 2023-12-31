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
    <br>
    <a href="/contact">Contact Us</a>
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
	<h1>Hello, {{.Name}}! Welcome to SIGNIFY.</h1>
    <h2>Signify is the world leader in lighting for professionals, consumers and lighting for the Internet of Things</h2>
    <p>Server: {{.ServerName}}</p>
</body>
</html>
`

// added contact template
const contactTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h2>Contact Us</h2>
    <form action="/contact" method="post">
        <label for="name">Name:</label>
        <input type="text" name="name" required>
        <br>
        <label for="email">Email:</label>
        <input type="email" name="email" required>
        <br>
        <label for="message">Message:</label>
        <textarea name="message" rows="4" required></textarea>
        <br>
        <input type="submit" value="Send">
    </form>
    <br>
    <a href="/">Back to Home</a>
</body>
</html>
`
