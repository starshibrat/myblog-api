<h1>MyBlog API</h1>
<h2>Description</h2>
<p>REST API for simple blog/social media/forum. The API is utilise go framework gin-gonic with mongodb as It's database.</p>
<h2>Implemented Features</h2>
<h3>User</h3>
<h4>Registration</h4>
<p>Implementing anti duplicate username and email for user, password hashing, and automated date creation</p>
<h4>Login</h4>
<p>Checking hash between user-input password and database</p>
<h4>Retrieve All User</h4>
<p>Using specialized struct/data type, the returned data is limited to _id, username, and creation date</p>
<h4>JWT Protected Middleware</h4>
<p>Sensitive API like post creation and deletion or user deletion will be protected by middleware first for token verification</p>

<h3>Posts</h3>
<h4>Creation</h4>
<p>Create new post that only send the content of the post</p>
<h4>Deletion</h4>
<p>Delete post</p>
<h4>Retrieve all posts</h4>
<p>Retrieving all post in the `posts` collection</p>
