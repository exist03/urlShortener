CREATE DATABASE url;
GRANT ALL PRIVILEGES ON DATABASE url TO postgres;
\c url
DROP TABLE IF EXISTS urls;
CREATE TABLE urls
(
    short VARCHAR(10) NOT NULL
        PRIMARY KEY,
    long  VARCHAR(255) NOT NULL
);