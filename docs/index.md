```
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <title>Web4App.io</title>
</head>
<body>
  <h1>🌐 Welcome to Web4App.io</h1>
  <p>This is your Go-powered full-stack server.</p>
  <script>
    fetch("/api/hello")
      .then(res => res.json())
      .then(data => {
        const msg = document.createElement("p");
        msg.innerText = "API says: " + data.message;
        document.body.appendChild(msg);
      });
  </script>
</body>
</html>
