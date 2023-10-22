package blog

import (
	"22nd_Oct_Antino/model"
	"database/sql"
	"errors"

	log "github.com/sirupsen/logrus"
)

type Dao struct {
	db                 *sql.DB
	getAllBlogPosts    *sql.Stmt
	getBlogPostById    *sql.Stmt
	addBlogPost        *sql.Stmt
	addAuthor          *sql.Stmt
	updateBlogPostById *sql.Stmt
	updateAuthorById   *sql.Stmt
	deleteBlogPostById *sql.Stmt
	deleteAuthorById   *sql.Stmt
	getAuthorId        *sql.Stmt
}

func NewDao(db *sql.DB) (*Dao, error) {

	dao := Dao{db: db}
	var err error

	dao.getAllBlogPosts, err = db.Prepare(getAllBlogPostsQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	dao.getBlogPostById, err = db.Prepare(getBlogPostByIdQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	dao.addBlogPost, err = db.Prepare(addBlogPostQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	dao.addAuthor, err = db.Prepare(addAuthorQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	dao.updateBlogPostById, err = db.Prepare(updateBlogPostByIdQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	dao.updateAuthorById, err = db.Prepare(updateAuthorByIdQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	dao.deleteBlogPostById, err = db.Prepare(deleteBlogPostByIdQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	dao.deleteAuthorById, err = db.Prepare(deleteAuthorByIdQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	dao.getAuthorId, err = db.Prepare(getAuthorIdQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &dao, nil
}

const addBlogPostQuery = `
		INSERT INTO BLOGPOSTS
        ( title, caption, body, note, datecreated, modifieddates, author_id )
        VALUES
        (?,?,?,?,?,?,?);
		SELECT SCOPE_IDENTITY() AS blogPostId;
		`
const addAuthorQuery = `
		INSERT INTO AUTHOR
		( firstname, lastname, fullname, dateofbirth, gender, house_no, house_name, street, 
			city, state, country, pincode  )
		VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?);
		SELECT SCOPE_IDENTITY() AS author_id;
		`

func (dao *Dao) AddBlogPost(blog_post model.BlogPost) (*int64, error) {

	var authorid int64
	author_err := dao.addAuthor.QueryRow(
		blog_post.First_Name,
		blog_post.Last_Name,
		blog_post.Full_Name,
		blog_post.Date_Of_Birth,
		blog_post.Gender,
		blog_post.House_No,
		blog_post.House_Name,
		blog_post.Street,
		blog_post.City,
		blog_post.State,
		blog_post.Country,
		blog_post.PinCode).Scan(&authorid)

	if author_err != nil {
		log.Error(author_err)
		return nil, author_err
	}

	var blogPostId int64
	post_err := dao.addBlogPost.QueryRow(
		blog_post.Title,
		blog_post.Caption,
		blog_post.Body,
		blog_post.Note,
		blog_post.CreatedDate,
		blog_post.ModifiedDate,
		authorid).Scan(&blogPostId)

	if post_err != nil {
		log.Error(post_err)
		return nil, post_err
	}

	return &blogPostId, nil
}

var ErrPostNotFound = errors.New("No Blog Post Found")

const updateBlogPostByIdQuery = `UPDATE [DBO].[BLOGPOSTS] SET 
		title = ?,
		caption = ?,
		body = ?,
		note = ?,
		datecreated= ?,
		modifieddates  = ?
	WHERE ID = ?`

const getAuthorIdQuery = ` SELECT author_id FROM [BLOGPOSTS] WHERE  ID = ?`

const updateAuthorByIdQuery = `UPDATE [DBO].[AUTHOR] SET 
		firstname = ?,
		lastname = ?,
		fullname = ?,
		dateofbirth = ?,
		gender = ?,
		house_no = ?,
		house_name = ?,
		street = ?, 
		city = ?,
		state = ?,
		country = ?,
		pincode = ?,
WHERE ID = ?`

func (dao *Dao) UpdateBlogPostById(blog_post model.BlogPost) (*int64, error) {

	authorID := dao.getAuthorId.QueryRow(blog_post.Id)
	var author_id *int64
	err := authorID.Scan(&author_id)
	if err == sql.ErrNoRows {
		return nil, ErrPostNotFound
	} else if err != nil {
		return nil, err
	}

	// Passing the constructed BlogPost struct for updation
	_, author_err := dao.updateBlogPostById.Exec(
		blog_post.Title,
		blog_post.Caption,
		blog_post.Body,
		blog_post.Note,
		blog_post.CreatedDate,
		blog_post.ModifiedDate,
		blog_post.Id)

	if author_err != nil {
		return nil, author_err
	}

	result, err := dao.updateAuthorById.Exec(
		blog_post.First_Name,
		blog_post.Last_Name,
		blog_post.Full_Name,
		blog_post.Date_Of_Birth,
		blog_post.Gender,
		blog_post.House_No,
		blog_post.House_Name,
		blog_post.Street,
		blog_post.City,
		blog_post.State,
		blog_post.Country,
		blog_post.PinCode,
		author_id)

	if err != nil {
		return nil, err
	}

	// Get the number of rows affected for confirmation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	} else if rowsAffected == 0 {
		return nil, ErrPostNotFound
	}

	return &rowsAffected, nil
}

const deleteBlogPostByIdQuery = "DELETE FROM BLOGPOSTS WHERE ID = ?"

const deleteAuthorByIdQuery = "DELETE FROM AUTHOR WHERE ID = ?"

func (dao *Dao) DeleteBlogPostById(blogPostId int64) error {

	authorID := dao.getAuthorId.QueryRow(blogPostId)
	var author_id *int64
	err := authorID.Scan(&author_id)
	if err == sql.ErrNoRows {
		return ErrPostNotFound
	} else if err != nil {
		return err
	}

	result, err := dao.deleteAuthorById.Exec(author_id)
	if err != nil {
		return err
	}

	result, err = dao.deleteBlogPostById.Exec(blogPostId)
	if err != nil {
		return err
	}

	// Get the number of rows affected for confirmation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return ErrPostNotFound
	}

	return nil
}

const getAllBlogPostsQuery = `
	SELECT b.Id, b.title, b.caption, b.body, b.note, b.datecreated, b.modifieddates, 
		a.Id AS author_id , a.firstname, a.lastname, a.fullname, a.dateofbirth, a.gender, a.house_no, 
		a.house_name, a.street, a.city, a.state, a.country, a.pincode   FROM BLOGPOSTS b , AUTHOR a 
		WHERE b.Id = a.Id
		GROUP BY b.id , b.title b.caption
`

const getAllBlogPostsCountQuery = `SELECT COUNT(*)CNT FROM BLOGPOSTS`

func (dao *Dao) GetAllBlogPosts(inputModel *model.SearchPostRequest) (*model.PaginatedPostsResponse, error) {

	var cnt *int64
	cntStmt, err := dao.db.Prepare(getAllBlogPostsCountQuery)
	if err != nil {
		return nil, err
	}
	err = cntStmt.QueryRow().Scan(&cnt)
	if err == sql.ErrNoRows {
		return nil, ErrPostNotFound
	} else if err != nil {
		return nil, err
	}

	stmt, err := dao.db.Prepare(getAllBlogPostsCountQuery)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var result *model.PaginatedPostsResponse
	var BlogPostList []*model.BlogPost

	for rows.Next() {
		var post model.BlogPost
		err := rows.Scan(&post.Id, &post.Title, &post.Caption,
			&post.Body, &post.CreatedDate, &post.ModifiedDate, &post.AuthorId,
			&post.First_Name, &post.Last_Name, &post.Full_Name, &post.Date_Of_Birth, &post.Gender,
			&post.House_No, &post.House_Name, &post.Street, &post.City, &post.State,
			&post.Country, &post.PinCode)

		if err != nil {
			return nil, err
		} else {
			BlogPostList = append(BlogPostList, &post)
		}
	}

	result = &model.PaginatedPostsResponse{
		Result:              BlogPostList,
		PageNumber:          int64(inputModel.PageNumber),
		PageSize:            int64(inputModel.PageSize),
		TotalRecords:        *cnt,
		SelectedRecordCount: len(BlogPostList),
		SortField:           inputModel.SortField,
		SortOrder:           inputModel.SortOrder,
	}

	return result, nil
}

const getBlogPostByIdQuery = `
		DECLARE @BLOGPOSTID = ?
		b.Id, b.title, b.caption, b.body, b.note, b.datecreated, b.modifieddates, 
		a.Id AS author_id , a.firstname, a.lastname, a.fullname, a.dateofbirth, a.gender, a.house_no, 
		a.house_name, a.street, a.city, a.state, a.country, a.pincode   FROM BLOGPOSTS b , AUTHOR a 
		WHERE a.Id = @BLOGPOSTID and b.Id = a.Id `

func (dao *Dao) GetBlogPostById(blogPostId int64) (*model.BlogPost, error) {

	// Execute getBlogPostByIdQuery
	row := dao.getBlogPostById.QueryRow(blogPostId)

	var post model.BlogPost

	// Scanning row into the BlogPost struct
	err := row.Scan(&post.Id, &post.Title, &post.Caption,
		&post.Body, &post.CreatedDate, &post.ModifiedDate, &post.AuthorId,
		&post.First_Name, &post.Last_Name, &post.Full_Name, &post.Date_Of_Birth, &post.Gender,
		&post.House_No, &post.House_Name, &post.Street, &post.City, &post.State,
		&post.Country, &post.PinCode)

	if err != nil {
		return nil, err
	}

	log.Error(err)

	return &post, nil
}
