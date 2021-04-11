CREATE TABLE event_types (
  code VARCHAR(24),
  name NVARCHAR(255),
  description TEXT,

  PRIMARY KEY (code)
);

INSERT INTO event_types VALUES 
('SMO', 'SMO', '');