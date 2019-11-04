<nav class="navbar navbar-inverse navbar-static-top">
    <div class="container-fluid">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false"
                aria-controls="navbar">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="">GO CRUD</a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
        {{.test}}
            {{if .UserID}}
                <ul class="nav navbar-nav navbar-right">
                    <li><a href="/profile">Profile</a></li>
                    <li><a href="/notepad">Notepad</a></li>
                    <li><a href="/logout">Logout</a></li>
                </ul>
            {{end}}
        </div>
    </div>
</nav>