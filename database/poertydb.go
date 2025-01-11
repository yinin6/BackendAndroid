package database

import (
	"BackendAndroid/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

func SaveToDB(db *sql.DB, apiResponse *models.APIResponse) error {
	// 插入主表数据
	query := `
		INSERT INTO api_data (id, content, popularity, token, ip_address, cache_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := db.Exec(query, apiResponse.Data.ID, apiResponse.Data.Content, apiResponse.Data.Popularity,
		apiResponse.Token, apiResponse.IPAddress, apiResponse.Data.CacheAt)
	if err != nil {
		return err
	}

	// 插入 origin 数据
	originQuery := `
		INSERT INTO origin (data_id, title, dynasty, author)
		VALUES (?, ?, ?, ?)
	`
	_, err = db.Exec(originQuery, apiResponse.Data.ID, apiResponse.Data.Origin.Title,
		apiResponse.Data.Origin.Dynasty, apiResponse.Data.Origin.Author)
	if err != nil {
		return err
	}

	// 插入 origin_content 数据
	for _, content := range apiResponse.Data.Origin.Content {
		contentQuery := `
			INSERT INTO origin_content (data_id, content)
			VALUES (?, ?)
		`
		_, err = db.Exec(contentQuery, apiResponse.Data.ID, content)
		if err != nil {
			return err
		}
	}

	// 插入 origin_translate 数据
	for _, translate := range apiResponse.Data.Origin.Translate {
		translateQuery := `
			INSERT INTO origin_translate (data_id, translate)
			VALUES (?, ?)
		`
		_, err = db.Exec(translateQuery, apiResponse.Data.ID, translate)
		if err != nil {
			return err
		}
	}

	// 插入 match_tags 数据
	for _, tag := range apiResponse.Data.MatchTags {
		tagQuery := `
			INSERT INTO match_tags (data_id, tag)
			VALUES (?, ?)
		`
		_, err = db.Exec(tagQuery, apiResponse.Data.ID, tag)
		if err != nil {
			return err
		}
	}

	return nil
}

func FetchFromDB(db *sql.DB, id string) (*models.APIResponse, error) {
	// 查询主表数据
	var data models.Data
	var token, ipAddress string
	err := db.QueryRow(`
		SELECT id, content, popularity, token, ip_address, cache_at
		FROM api_data
		WHERE id = ?
	`, id).Scan(&data.ID, &data.Content, &data.Popularity, &token, &ipAddress, &data.CacheAt)
	if err != nil {
		return nil, err
	}

	// 查询 origin 数据
	var origin models.Origin
	err = db.QueryRow(`
		SELECT title, dynasty, author
		FROM origin
		WHERE data_id = ?
	`, id).Scan(&origin.Title, &origin.Dynasty, &origin.Author)
	if err != nil {
		return nil, err
	}

	// 查询 origin_content 数据
	rows, err := db.Query(`
		SELECT content
		FROM origin_content
		WHERE data_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var content string
		err = rows.Scan(&content)
		if err != nil {
			return nil, err
		}
		origin.Content = append(origin.Content, content)
	}

	// 查询 origin_translate 数据
	rows, err = db.Query(`
		SELECT translate
		FROM origin_translate
		WHERE data_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var translate string
		err = rows.Scan(&translate)
		if err != nil {
			return nil, err
		}
		origin.Translate = append(origin.Translate, translate)
	}

	// 查询 match_tags 数据
	rows, err = db.Query(`
		SELECT tag
		FROM match_tags
		WHERE data_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matchTags []string
	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			return nil, err
		}
		matchTags = append(matchTags, tag)
	}

	// 组装数据
	data.Origin = origin
	data.MatchTags = matchTags

	apiResponse := &models.APIResponse{
		Status:    "success",
		Data:      data,
		Token:     token,
		IPAddress: ipAddress,
	}

	return apiResponse, nil
}

// FetchRandomFromDB 从数据库中随机获取 N 条数据
func FetchRandomFromDB(db *sql.DB, n int) ([]*models.APIResponse, error) {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 查询主表数据，随机获取 N 条
	rows, err := db.Query(`
		SELECT id, content, popularity, token, ip_address, cache_at
		FROM api_data
		ORDER BY RAND()
		LIMIT ?
	`, n)
	if err != nil {
		return nil, fmt.Errorf("failed to query api_data: %v", err)
	}
	defer rows.Close()

	var responses []*models.APIResponse

	// 遍历查询结果
	for rows.Next() {
		var data models.Data
		var token, ipAddress string

		// 扫描主表数据
		err := rows.Scan(&data.ID, &data.Content, &data.Popularity, &token, &ipAddress, &data.CacheAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan api_data: %v", err)
		}

		// 查询 origin 数据
		var origin models.Origin
		err = db.QueryRow(`
			SELECT title, dynasty, author
			FROM origin
			WHERE data_id = ?
		`, data.ID).Scan(&origin.Title, &origin.Dynasty, &origin.Author)
		if err != nil {
			return nil, fmt.Errorf("failed to query origin: %v", err)
		}

		// 查询 origin_content 数据
		contentRows, err := db.Query(`
			SELECT content
			FROM origin_content
			WHERE data_id = ?
		`, data.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to query origin_content: %v", err)
		}
		defer contentRows.Close()

		for contentRows.Next() {
			var content string
			err = contentRows.Scan(&content)
			if err != nil {
				return nil, fmt.Errorf("failed to scan origin_content: %v", err)
			}
			origin.Content = append(origin.Content, content)
		}

		// 查询 origin_translate 数据
		translateRows, err := db.Query(`
			SELECT translate
			FROM origin_translate
			WHERE data_id = ?
		`, data.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to query origin_translate: %v", err)
		}
		defer translateRows.Close()

		for translateRows.Next() {
			var translate string
			err = translateRows.Scan(&translate)
			if err != nil {
				return nil, fmt.Errorf("failed to scan origin_translate: %v", err)
			}
			origin.Translate = append(origin.Translate, translate)
		}

		// 查询 match_tags 数据
		tagRows, err := db.Query(`
			SELECT tag
			FROM match_tags
			WHERE data_id = ?
		`, data.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to query match_tags: %v", err)
		}
		defer tagRows.Close()

		var matchTags []string
		for tagRows.Next() {
			var tag string
			err = tagRows.Scan(&tag)
			if err != nil {
				return nil, fmt.Errorf("failed to scan match_tags: %v", err)
			}
			matchTags = append(matchTags, tag)
		}

		// 组装数据
		data.Origin = origin
		data.MatchTags = matchTags

		// 创建 APIResponse 对象
		response := &models.APIResponse{
			Status:    "success",
			Data:      data,
			Token:     token,
			IPAddress: ipAddress,
		}

		// 添加到结果集
		responses = append(responses, response)
	}

	// 检查遍历过程中是否有错误
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return responses, nil
}
