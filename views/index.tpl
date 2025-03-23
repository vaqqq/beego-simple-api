<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>BeeGo Simple API</title>
  <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
	<div class="card">
    <div class="header">
    	<img src="/static/img/beego.png" alt="Beego Icon" class="logo">
    	<h1>BeeGo simple API</h1>
      </div>
    <p>Select an endpoint to test:</p>

    <div class="button-group">
      <button id="btnHello">GET /hello</button>
	    <button id="btnTime">GET /time</button>
	    <button id="btnPost">Send POST</button>
    </div>

    <h3>POST /post with message:</h3>
    <input type="text" id="postMessage" placeholder="Enter your message" style="padding: 5px; width: 60%;">

		<h3>Response:</h3>
		<pre id="output">Click a button to load response...</pre>

		<hr>
    <p>Developed with:</p>

    <ul>
    	{{range .Modules}}
    		<li><a href="{{.Link}}" target="_blank">{{.Name}}</a></li>
    	{{end}}
    </ul>

		<p>GitHub: <a href="{{.GitHub}}" target="_blank">{{.GitHub}}</a></p>
	</div>
  <script src="/static/js/script.js"></script>
</body>
</html>
