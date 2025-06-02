package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"product-service/internal/entities"
	"time"

	"github.com/redis/go-redis/v9"
)

type ProductRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewProductRepository(db *sql.DB, rdb *redis.Client) *ProductRepository {
	return &ProductRepository{
		db:  db,
		rdb: rdb,
	}
}

func (repo *ProductRepository) CreateProduct(ctx context.Context, pr *entities.Product) error {
	const op = "storage.createproduct"
	query := `
		INSERT INTO products(name, description, price, image_url, seller_id, created_at)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	if pr.CreatedAt.IsZero() {
		pr.CreatedAt = time.Now()
	}

	err := repo.db.QueryRowContext(ctx, query, pr.Name, pr.Description, pr.Price, pr.ImgUrl, pr.SellerID, pr.CreatedAt).Scan(&pr.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (repo *ProductRepository) GetProducts(ctx context.Context) ([]*entities.Product, error) {
	const op = "storage.getproducts"
	query := `
		SELECT * FROM products
	`
	rows, err := repo.db.QueryContext(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	var products []*entities.Product

	for rows.Next() {
		var pr entities.Product
		if err = rows.Scan(
			&pr.ID,
			&pr.Name,
			&pr.Description,
			&pr.Price,
			&pr.ImgUrl,
			&pr.SellerID,
			&pr.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		products = append(products, &pr)
	}

	return products, nil
}

func (repo *ProductRepository) GetProductById(ctx context.Context, id int64) (*entities.Product, error) {
	const op = "storage.getproductbyid"
	key := fmt.Sprintf("product:%d", id)
	val, err := repo.rdb.Get(ctx, key).Result()
	if err == nil {
		var product entities.Product
		if err = json.Unmarshal([]byte(val), &product); err == nil {
			return &product, nil
		}
	}

	query := `
		SELECT * FROM products
		WHERE id = $1
	`
	var pr entities.Product
	err = repo.db.QueryRowContext(ctx, query, id).Scan(
		&pr.ID,
		&pr.Name,
		&pr.Description,
		&pr.Price,
		&pr.ImgUrl,
		&pr.SellerID,
		&pr.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	data, _ := json.Marshal(pr)
	repo.rdb.Set(ctx, key, data, time.Hour)
	return &pr, nil
}

func (repo *ProductRepository) DeleteProduct(ctx context.Context, id int64) error {
	const op = "storage.deleteproduct"
	query := `DELETE FROM products WHERE id = $1`
	_, err := repo.db.ExecContext(ctx, query, id)
	if err == nil {
		key := fmt.Sprintf("product:%d", id)
		repo.rdb.Del(ctx, key)
	}
	return fmt.Errorf("%s: %w", op, err)
}

func (repo *ProductRepository) UpdateProduct(ctx context.Context, id int64, u *entities.Product) error {
	const op = "storage.updateproduct"

	_, err := repo.GetProductById(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	query := `
		UPDATE products SET
		name = $1,
		description = $2,
		price = $3,
		image_url = $4,
		WHERE id = $5
	`
	_, err = repo.db.ExecContext(ctx, query, u.Name, u.Description, u.Price, u.ImgUrl, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
