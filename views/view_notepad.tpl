<div class="container">
	<div class="well">
		<div class="media">
			<div class="media-body">
				<h4 class="media-heading">
					<strong>Title will come soon</strong>
				</h4>
				<p class="text-right">By 
					<strong>{{.UserName}}</strong>
				</p>
				<p>{{.Note.Comments}}</p>
				<ul class="list-inline list-unstyled">
					<li>
						<span>
							<i class="glyphicon glyphicon-calendar"></i> 
              {{.Note.CreatedAt.Format "02-01-2006 15:04:05"}} 
						</span>
					</li>
					<li>
						<span>
							<i class="glyphicon glyphicon-time"></i>        {{.Note.UpdatedAt.Format "02-01-2006 15:04:05"}} 
						</span>
					</li>
				</ul>
			</div>
		</div>
	</div>
	<a href="/notepad" class="btn btn-danger"><span aria-hidden="true"></span><span><strong>Back</strong></span></a>
</div>