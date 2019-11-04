<div class="container">
	<div class="row">
            <div class="col-md-12">
                <div class="panel panel-default">
                    <div class="panel-heading">
                        <h3 class="panel-title">Welcome {{.UserName}}</h3>
                    </div>
                    <div class="panel-body">
                        <div class="row">
                            <div class="col-md-3 col-lg-3 hidden-xs hidden-sm">
                                <img class="img-circle" src="/static/img/no-img.png" width="100" height="100" alt="User Pic">
                            </div>
                            <div class=" col-md-12">
                                <table class="table">
                                    <tbody>
                                    <tr>
                                        <td class="table-no-border">First Name</td>
                                        <td class="table-no-border">{{.User.FirstName}}</td>
                                    </tr>
                                    <tr>
                                        <td>Last Name</td>
                                        <td>{{.User.LastName}}</td>
                                    </tr>
                                    <tr>
                                        <td>Email</td>
                                        <td>{{.User.Email}}</td>
                                    </tr>
                                    <tr>
                                        <td>Phone</td>
                                        <td>{{.User.Phone}}</td>
                                    </tr>
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                    <!-- <div class="panel-footer">
                        <button class="btn btn-sm btn-primary" type="button" data-toggle="tooltip" data-original-title="Send message to user"><i class="glyphicon glyphicon-envelope"></i></button>
                        <span class="pull-right">
                            <button class="btn btn-sm btn-warning" type="button" data-toggle="tooltip" data-original-title="Edit this user"><i class="glyphicon glyphicon-edit"></i></button>
                            <button class="btn btn-sm btn-danger" type="button" data-toggle="tooltip" data-original-title="Remove this user"><i class="glyphicon glyphicon-remove"></i></button>
                        </span>
                    </div> -->
                </div>
            </div>
        </div>
</div>
    