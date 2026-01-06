from django.db.models.expressions import result
from flask import Flask, request, jsonify, render_template, abort
from flask_sqlalchemy import SQLAlchemy
from sqlalchemy import func
from models import db
from models import LessonPlan
from markupsafe import Markup
import markdown as md
import requests
import os
import re
import json
from LLM_api import (
    tongyi_generate_objectives,
    tongyi_generate_key,
    tongyi_generate_difficult,
    tongyi_generate_content,
    tongyi_generate_ideological,
    tongyi_generate_reflection,
)

# ========== Flask & DB 基础配置 ==========
app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = \
    'mysql+pymysql://root:123456@127.0.0.1:3306/studyonline?charset=utf8mb4'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db.init_app(app)

# 第一次运行前创建表
with app.app_context():
    db.create_all()

# ========== 业务辅助函数 ==========
def check_objectives(objectives, content):
    return all(obj.strip() in content for obj in objectives.split('\n') if obj.strip())

def check_key_difficult(key_points, difficult_points):
    return key_points.strip() != difficult_points.strip()

def auto_add_ideological_points(content):
    if '创新' in content:
        return '创新精神'
    if '团队' in content:
        return '团队协作'
    return '社会主义核心价值观'

def _get_data_from_request():
    """统一兼容 application/json 与 form 表单"""
    if request.content_type and 'application/json' in request.content_type:
        try:
            return request.get_json(silent=True) or {}
        except Exception:
            return {}
    return request.form or {}

# ========== 路由（接口） ==========

# 生成教案（支持分步 AI 生成；GET 渲染页面）
@app.route('/lesson/plan/generate', methods=['POST'])
def generate_lessonplan():
    data = _get_data_from_request()
    theme = data.get('theme', '') or ''
    duration = data.get('duration', '') or ''
    try:
        duration_int = int(duration)
    except Exception:
        duration_int = 45

    step = data.get('step')

    # 分步 AI 生成：仅返回 JSON，不落库
    if step != "":
        try:
            if (not step) or step == 'objectives':
                ai_content = tongyi_generate_objectives(theme, duration_int)
                if not ai_content or len(ai_content.strip()) == 0:
                    return jsonify({
                        'success': False,
                        'msg': '教学目标生成失败，请检查API密钥或主题内容',
                        'error': ai_content or '未知错误',
                        'objectives': ''
                    })
                cleaned = ai_content.strip()
                if "教学目标" in cleaned:
                    cleaned = cleaned.split("教学目标")[-1].strip()
                if "示例" in cleaned:
                    cleaned = cleaned.split("示例")[0].strip()
                return jsonify({'success': True, 'objectives': cleaned, 'raw': ai_content.strip()})

            elif step == 'key_difficult':
                objectives = data.get('objectives', '') or ''
                ai_content = tongyi_generate_key(theme, duration_int, objectives)
                if ai_content and len(ai_content.strip()) > 0:
                    key_points = ai_content.strip()
                ai_content = tongyi_generate_difficult(theme, duration_int, objectives)
                if ai_content and len(ai_content.strip()) > 0:
                    difficult_points = ai_content.strip()
                if not key_points and not difficult_points:
                    return jsonify({'success': False, 'msg': ai_content or 'AI生成失败', 'key_points': '', 'difficult_points': ''})
                return jsonify({'success': True, 'key_points': key_points, 'difficult_points': difficult_points, 'raw': (ai_content or '').strip()})

            elif step == 'content':
                objectives = data.get('objectives', '') or ''
                key_points = data.get('key_points', '') or ''
                difficult_points = data.get('difficult_points', '') or ''
                ai_content = tongyi_generate_content(theme, duration_int, objectives, key_points, difficult_points)
                if not ai_content or len((ai_content or '').strip()) == 0:
                    return jsonify({'success': False, 'msg': ai_content or 'AI生成失败', 'content': ''})
                return jsonify({'success': True, 'content': ai_content.strip()})

            elif step == 'ideological':
                objectives = data.get('objectives', '') or ''
                key_points = data.get('key_points', '') or ''
                difficult_points = data.get('difficult_points', '') or ''
                content = data.get('content', '') or ''
                ai_content = tongyi_generate_ideological(theme, duration_int, objectives, key_points, difficult_points, content)
                if ai_content and len(ai_content.strip()) > 0:
                    return jsonify({'success': True, 'ideological_points': ai_content.strip()})
                return jsonify({'success': False, 'msg': ai_content or 'AI生成失败', 'ideological_points': ''})

            else:
                return jsonify({'success': False, 'msg': '未指定生成步骤'})
        except Exception as e:
            return jsonify({'success': False, 'msg': f'AI生成异常: {e}'})

    try:
        # 步骤1: 生成教学目标
        objectives = ''
        try:
            ai_content = tongyi_generate_objectives(theme, duration_int)
            print(ai_content)
            if ai_content and len(ai_content.strip()) > 0:
                objectives = ai_content.strip()
        except Exception as e:
            print(f"生成教学目标失败: {e}")

        # 步骤2: 生成重点难点
        key_points = ''
        difficult_points = ''
        if objectives:
            try:
                ai_content = tongyi_generate_key(theme, duration_int, objectives)
                if ai_content and len(ai_content.strip()) > 0:
                    key_points = ai_content.strip()
                ai_content = tongyi_generate_difficult(theme, duration_int, objectives)
                if ai_content and len(ai_content.strip()) > 0:
                    difficult_points = ai_content.strip()
            except Exception as e:
                print(f"生成重点难点失败: {e}")

        # 步骤3: 生成教学内容
        content = ''
        if objectives or key_points or difficult_points:
            try:
                ai_content = tongyi_generate_content(theme, duration_int, objectives, key_points, difficult_points)
                if ai_content and len((ai_content or '').strip()) > 0:
                    content = ai_content.strip()
            except Exception as e:
                print(f"生成教学内容失败: {e}")

        # 步骤4: 生成思政要点
        ideological_points = ''
        if content:
            try:
                ai_content = tongyi_generate_ideological(theme, duration_int, objectives, key_points, difficult_points,
                                                         content)
                if ai_content and len(ai_content.strip()) > 0:
                    ideological_points = ai_content.strip()
            except Exception as e:
                print(f"生成思政要点失败: {e}")


        # 如果没有生成思政要点，尝试自动生成
        if not ideological_points and content:
            ideological_points = auto_add_ideological_points(content)

        # 保存到数据库
        plan = LessonPlan(
            title=theme,
            duration=duration_int,
            objectives=objectives,
            key_points=key_points,
            difficult_points=difficult_points,
            content=content,
            ideological_points=ideological_points
        )
        db.session.add(plan)
        db.session.commit()

        # 返回完整结果
        return jsonify({
            'success': True,
            'msg': '教案已自动生成并保存',
            'id': plan.id,
            'data': {
                'objectives': objectives,
                'key_points': key_points,
                'difficult_points': difficult_points,
                'content': content,
                'ideological_points': ideological_points
            },
            'auto_completed': True
        })

    except Exception as e:
        return jsonify({'success': False, 'msg': f'自动生成教案异常: {e}', 'auto_completed': True})
    # # 正常表单提交：保存教案
    # objectives = data.get('objectives', '') or ''
    # key_points = data.get('key_points', '') or ''
    # difficult_points = data.get('difficult_points', '') or ''
    # content = data.get('content', '') or ''
    # ideological_points = (data.get('ideological', '') or '').strip()
    # if not ideological_points:
    #     ideological_points = auto_add_ideological_points(content)
    #
    # plan = LessonPlan(
    #     title=theme,
    #     duration=duration_int,
    #     objectives=objectives,
    #     key_points=key_points,
    #     difficult_points=difficult_points,
    #     content=content,
    #     ideological_points=ideological_points
    # )
    # db.session.add(plan)
    # db.session.commit()
    # return jsonify({'success': True, 'msg': '教案已保存', 'id': plan.id})

# 教案评价（GET）
@app.route('/lessonplan/<int:pk>/evaluate', methods=['GET'])
def evaluate_lessonplan(pk):
    plan = LessonPlan.query.get_or_404(pk)
    result = {
        'objectives_achieved': check_objectives(plan.objectives, plan.content),
        'key_difficult_distinct': check_key_difficult(plan.key_points, plan.difficult_points),
        'ideological_points': plan.ideological_points or auto_add_ideological_points(plan.content)
    }
    return jsonify(result)

# 教学反思（仅支持 POST）
@app.route('/lessonplan/<int:pk>/reflect', methods=['POST'])
def reflect_lessonplan(pk):
    plan = LessonPlan.query.get_or_404(pk)
    reflection = tongyi_generate_reflection(
        plan.title,
        plan.duration,
        plan.objectives,
        plan.key_points,
        plan.difficult_points,
        plan.content,
        plan.ideological_points
    )
    return jsonify({'success': True, 'reflection': reflection})


if __name__ == '__main__':
    # 默认开发模式运行
    app.run(host='0.0.0.0', port=int(os.getenv('PORT', 12010)), debug=True)