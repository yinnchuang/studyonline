from flask_sqlalchemy import SQLAlchemy


db = SQLAlchemy()
class LessonPlan(db.Model):
    __tablename__ = 'lesson_plans'
    id = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String(255), default='')
    duration = db.Column(db.Integer, default=45)
    objectives = db.Column(db.Text, default='')
    key_points = db.Column(db.Text, default='')
    difficult_points = db.Column(db.Text, default='')
    content = db.Column(db.Text, default='')
    ideological_points = db.Column(db.Text, default='')