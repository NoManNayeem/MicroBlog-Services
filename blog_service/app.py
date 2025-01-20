from flask import Flask, jsonify
from flask_restful import Api
from flask_jwt_extended import JWTManager
from models import db_init
from resources import BlogList, BlogResource
import os
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv()

# Initialize Flask app
app = Flask(__name__)
api = Api(app)

# Configure the SQLite3 database
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///database.db'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

# Configure JWT secret key (must match DRF secret key)
app.config['JWT_SECRET_KEY'] = os.getenv("JWT_SECRET_KEY")  # Use .env value
app.config['JWT_IDENTITY_CLAIM'] = 'user_id'  # Use 'user_id' as the identity claim

# Initialize database and JWT
db_init(app)
jwt = JWTManager(app)

# Add resources to the API
api.add_resource(BlogList, '/blogs')               # Protected blogs list
api.add_resource(BlogResource, '/blogs/<int:blog_id>')  # Protected single blog

@app.route('/')
def home():
    return jsonify({"message": "Welcome to the Blog Microservice!"})

if __name__ == '__main__':
    app.run(debug=True)
