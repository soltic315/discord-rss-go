CREATE TABLE feeds (
    feed_id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    url VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (url)
);

CREATE TABLE articles (
    article_id SERIAL PRIMARY KEY,
    feed_id INTEGER NOT NULL,
    title VARCHAR(100) NOT NULL,
    link VARCHAR(100) NOT NULL,
    published_at TIMESTAMP NOT NULL,
    need_notify BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (feed_id) REFERENCES feeds(feed_id) ON DELETE CASCADE,
    UNIQUE (feed_id, published_at)
);

CREATE TABLE subscriptions (
    subscription_id SERIAL PRIMARY KEY,
    feed_id INTEGER NOT NULL,
    channel_id VARCHAR(100) NOT NULL,
    created_by VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (feed_id) REFERENCES feeds(feed_id),
    UNIQUE (feed_id, channel_id)
);
