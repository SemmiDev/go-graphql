CREATE TABLE IF NOT EXISTS Authors(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    FirstName VARCHAR(255) NOT NULL,
    LastName VARCHAR(255) NOT NULL,
    PRIMARY KEY  (ID)
)