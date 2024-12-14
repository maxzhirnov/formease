package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/maxzhirnov/formease/internal/models"
)

type FormRepository struct {
	db *sql.DB
}

func NewFormRepository(db *sql.DB) *FormRepository {
	return &FormRepository{db: db}
}

func (r *FormRepository) GetForm(id int) (*models.Form, error) {
	var form models.Form
	var thankYouMessageJSON, questionsJSON []byte

	err := r.db.QueryRow(`
        SELECT f.id, f.name, f.theme, f.floating_shapes, 
               json_build_object(
                   'title', tm.title,
                   'subtitle', tm.subtitle,
                   'icon', tm.icon,
                   'button', json_build_object(
                       'text', tm.button_text,
                       'url', tm.button_url,
                       'newTab', tm.button_new_tab
                   )
               ) as thank_you_message,
               json_agg(
                   json_build_object(
                       'id', q.id,
                       'type', q.type,
                       'question', q.question,
                       'subtext', q.subtext,
                       'image', q.image,
                       'gradient', q.gradient,
                       'bgColor', q.bg_color,
                       'inputType', q.input_type,
                       'placeholder', q.placeholder,
                       'validation', q.validation,
                       'maxSelections', q.max_selections,
                       'options', (
                           SELECT json_agg(
                               json_build_object(
                                   'text', qo.text,
                                   'icon', qo.icon,
                                   'color', qo.color
                               )
                           )
                           FROM question_options qo
                           WHERE qo.question_id = q.id
                       )
                   )
               ) as questions
        FROM forms f
        LEFT JOIN thank_you_messages tm ON f.thank_you_message_id = tm.id
        LEFT JOIN questions q ON q.form_id = f.id
        WHERE f.id = $1
        GROUP BY f.id, tm.id`,
		id,
	).Scan(&form.ID, &form.Name, &form.Theme, &form.FloatingShapes,
		&thankYouMessageJSON, &questionsJSON)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(thankYouMessageJSON, &form.ThankYouMessage); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(questionsJSON, &form.Questions); err != nil {
		return nil, err
	}

	return &form, nil
}

// Add other necessary methods (CreateForm, UpdateForm, DeleteForm)
