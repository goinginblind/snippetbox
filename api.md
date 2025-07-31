| Route Pattern | Handler |              Action |
|---------------|---------|---------------------|
|`GET /{$}`       |`home`     |Display the home page|
|`GET /snippet/view/{id}` | `snippetView` |Display a specific snippet|
|`GET /snippet/create` | `snippetCreate` |Display a form for creating a new snippet|
|`POST /snippet/create` | `snippetCreatePost` |Save a new snippet|
