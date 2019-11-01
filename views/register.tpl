<div class="container">
    {{if .Errors}}
        <div class="alert alert-danger">
            {{range $value := .Errors}}
            <li> {{$value}} </li>
            {{end}}
        </div>
    {{end}}
    <form method="post" action="/register">
    <input autocomplete="false" name="hidden" type="text" style="display:none;">
        <div class="form-group">
            <label for="first_name">First Name</label>
            <div><input type="text" class="form-control" name="first_name" id="first_name" placeholder="First Name" value="{{.firstName}}" /></div>
        </div>

        <div class="form-group">
            <label for="last_name">Last Name</label>
            <div><input type="text" class="form-control" name="last_name" id="last_name" placeholder="Last Name" value="{{.lastName}}"/></div>
        </div>

        <div class="form-group">
            <label for="last_name">Email</label>
            <div><input type="text" class="form-control" name="email" id="email" placeholder="Email" value="{{.email}}"/></div>
        </div>

        <div class="form-group">
            <label for="last_name">Password</label>
            <div><input type="password" class="form-control" name="password" id="password" placeholder="Password" value="{{.password}}"/></div>
        </div>

        <div class="form-group">
            <label for="last_name">Confirm Password</label>
            <div><input type="password" class="form-control" name="confirm_password" id="confirm_password" placeholder="Confirm Password" value="{{.confirm_password}}"/></div>
        </div>

        <div class="form-group">
            <label for="last_name">Phone</label>
            <div><input type="text" class="form-control" name="phone" id="phone" placeholder="Phone" value="{{.phone}}"/></div>
        </div>

        <input type="submit" class="btn btn-primary" value="Register" />
    </form>
</div>