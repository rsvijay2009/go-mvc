<div class="container">
    {{if .Errors}}
	    <div class="alert alert-danger">
            {{range $value := .Errors}}
		        <li> {{$value}} </li>
            {{end}}
	    </div>
    {{end}}
	<form method="post" action="/notepad">
    <input type="hidden" name="noteId" value="{{.Note.Id}}">
		<div class="form-group">
			<label for="comments">{{if .Note}} Update {{else}} Create {{end}} your note</label>
			<textarea class="form-control" id="comments" name="comments" rows="7" placeholder="Enter your notes" required>{{if .Note}}{.Note.Comments}}{{end}}</textarea>
		</div>
		<input type="submit" class="btn btn-primary" 
        {{if .Note}} value="Update" {{else}} value="Create" {{end}}/>
        {{if .Note}}
		    <a href="/notepad" class="btn btn-danger"><span aria-hidden="true"></span><span><strong>Back</strong></span></a>
        {{end}}
	</form>
{{if .Notes}}
	<div class="row">
		<div class="panel widget">
			<div class="panel-body">
				<ul class="list-group">
                    {{range $note := .Notes}}
					<li class="list-group-item">
						<div class="row">
							<div class="col-xs-2 col-md-1">
								<img src="http://placehold.it/80" class="img-circle img-responsive" alt="" />
							</div>
							<div class="col-xs-10 col-md-11">
								<div>{{$note.Comments}}</div>
								<div class="text-info mt-3">Created on {{ $note.CreatedAt.Format "Jan 02, 2006" }}</div>
								<div class="action mt-3">
									<a href="notepad/edit/{{$note.Id}}" class="btn btn-primary"><span class="glyphicon glyphicon-edit" aria-hidden="true"></span><span><strong>Edit</strong></span>
									</a>
									<a href="notepad/view/{{$note.Id}}"  class="btn btn-success"><span class="glyphicon glyphicon-eye-open" aria-hidden="true"></span><span> <strong>View</strong></span>
									</a>
									<a href="javascript:void(0);" onclick="deleteNote({{$note.Id}})"class="btn btn-danger"><span class="glyphicon glyphicon-remove" aria-hidden="true"></span><span> <strong>Delete</strong></span>
									</a>
								</div>
							</div>
						</div>
					</li>
                    {{end}}
				</ul>
			</div>
		</div>
	</div>
{{end}}
</div>