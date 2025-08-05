| Route Pattern             | Handler               |Action                                     | 
|---------------------------|-----------------------|-------------------------------------------|
|`GET /{$}`                 |`home`                 |Display the home page                      | 
|`GET /snippet/view/{id}`   |`snippetView`          |Display a specific snippet                 |
|`GET /snippet/create`      |`snippetCreate`        |Display a form for creating a new snippet  |
|`POST /snippet/create`     |`snippetCreatePost`    |Save a new snippet                         |
|`GET /static/`             |`http.FileServer`      |Serve a specific static file               |
|`GET /user/signup`         |`userSignup`           |Display a form for user signup             |
|`POST /user/signup`        |`userSignupPost`       |Process signup form and create user        |
|`GET /user/login`          |`userLogin`            |Display a form for user login              |
|`POST /user/login`         |`userLoginPost`        |Process login form and authenticate user   |
|`POST /user/logout`        |`userLogoutPost`       |Log out the current user                   |