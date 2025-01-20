from flask_restful import Resource, reqparse
from flask_jwt_extended import jwt_required, get_jwt
from models import Blog, db

# Request parsers
blog_parser = reqparse.RequestParser()
blog_parser.add_argument('title', type=str, required=True, help='Title is required')
blog_parser.add_argument('content', type=str, required=True, help='Content is required')

# Blog list resource
class BlogList(Resource):
    @jwt_required()  # Require a valid JWT token
    def get(self):
        # Extract user_id from the JWT token
        current_user_id = get_jwt().get("user_id")
        if not current_user_id:
            return {"message": "Invalid token: Missing user_id"}, 400

        blogs = Blog.query.all()
        return [
            {"id": blog.id, "title": blog.title, "content": blog.content, "author": blog.author}
            for blog in blogs
        ], 200

    @jwt_required()  # Require a valid JWT token
    def post(self):
        data = blog_parser.parse_args()

        # Extract user_id from the JWT token
        current_user_id = get_jwt().get("user_id")
        if not current_user_id:
            return {"message": "Invalid token: Missing user_id"}, 400

        new_blog = Blog(title=data['title'], content=data['content'], author=current_user_id)
        db.session.add(new_blog)
        db.session.commit()
        return {
            "message": "Blog created successfully!",
            "blog": {
                "id": new_blog.id,
                "title": new_blog.title,
                "content": new_blog.content,
                "author": new_blog.author,
            },
        }, 201


# Single blog resource
class BlogResource(Resource):
    @jwt_required()  # Require a valid JWT token
    def get(self, blog_id):
        blog = Blog.query.get(blog_id)
        if not blog:
            return {"message": "Blog not found"}, 404
        return {"id": blog.id, "title": blog.title, "content": blog.content, "author": blog.author}, 200

    @jwt_required()  # Require a valid JWT token
    def put(self, blog_id):
        blog = Blog.query.get(blog_id)
        if not blog:
            return {"message": "Blog not found"}, 404

        # Extract user_id from the JWT token
        current_user_id = get_jwt().get("user_id")
        if not current_user_id:
            return {"message": "Invalid token: Missing user_id"}, 400

        data = blog_parser.parse_args()
        blog.title = data['title']
        blog.content = data['content']
        blog.author = current_user_id  # Update the author based on the current token
        db.session.commit()
        return {
            "message": "Blog updated successfully!",
            "blog": {"id": blog.id, "title": blog.title, "content": blog.content, "author": blog.author},
        }, 200

    @jwt_required()  # Require a valid JWT token
    def delete(self, blog_id):
        blog = Blog.query.get(blog_id)
        if not blog:
            return {"message": "Blog not found"}, 404
        db.session.delete(blog)
        db.session.commit()
        return {"message": "Blog deleted successfully!"}, 200
