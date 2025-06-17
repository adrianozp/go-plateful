CREATE TABLE places (
    id          VARCHAR(36)     PRIMARY KEY,
    name        VARCHAR(255)    NOT NULL,
    address     VARCHAR(500),
    phone       VARCHAR(50),
    email       VARCHAR(255),
    location    VARCHAR(255),
    category    VARCHAR(100),
    description TEXT,
    image       VARCHAR(2083),
    rating      DOUBLE,
    reviews     INT,
    tags        JSON,
    created_at  DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    updated_at  DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)
);

CREATE TABLE reviews (
	id  VARCHAR(36) PRIMARY KEY,
	user_id VARCHAR(36) NOT NULL,
	place_id VARCHAR(36) NOT NULL,
	rating      DOUBLE NOT NULL,
	content TEXT,
	created_at  DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    updated_at  DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    status VARCHAR(100) DEFAULT 'inactive'
)