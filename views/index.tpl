<div class="container">
    {{if .ErrorMsg}}
        <div class="alert alert-danger">
                {{.ErrorMsg}}
        </div>
    {{end}}
	<form method="post">
		<div class="form-group">
			<label for="email">Email Address</label>
			<div>
				<input type="email" class="form-control" name="email" id="email" maxlength="48" placeholder="Email" />
			</div>
		</div>
		<div class="form-group">
			<label for="password">Password</label>
			<div>
				<input type="password" class="form-control" name="password" id="password" maxlength="48" placeholder="Password" />
			</div>
		</div>
		<input type="submit" class="btn btn-primary" value="Login" />  Don't have account? 
		<a href="/register">register</a> here
	</form>
</div>