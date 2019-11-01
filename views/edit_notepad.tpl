<div class="container">
    {{if .Errors}}
	    <div class="alert alert-danger">
            {{range $value := .Errors}}
		        <li> {{$value}} </li>
            {{end}}
	    </div>
    {{end}}
	<form method="post" action="/notepad/edit/{{.Note.Id}}">
		<div class="form-group">
			<label for="comments">Create your note</label>
			<textarea class="form-control" id="comments" name="comments" rows="7" placeholder="Enter your notes" required>{{if .Note}}{{.Note.Comments}}{{end}}
            </textarea>
		</div>
		<input type="submit" class="btn btn-primary" value="Update"/>
		    <a href="/notepad" class="btn btn-danger"><span aria-hidden="true"></span><span><strong>Back</strong></span></a>
	</form>
</div>